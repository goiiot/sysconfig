package shell

import (
	"bufio"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"sync"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kr/pty"
	"go.uber.org/zap"
	"github.com/goiiot/sysconfig/cmd/server/conf"
	"github.com/goiiot/sysconfig/impl/log"
	"github.com/goiiot/sysconfig/impl/service/utils"
)

var (
	defaultShell string
)

func InitServiceShell(v1 *gin.RouterGroup, config *conf.ServiceShell) {
	if !config.Enabled {
		return
	}
	defaultShell = config.DefaultShell

	v1.GET("/shell", handleShellList)
	v1.POST("/shell", handleShellCreate)
	v1.DELETE("/shell/:id", handleShellClose)
	v1.POST("/shell/:id/size", handleShellSize)
	v1.GET("/shell/:id", handleShellSession)
}

var (
	up        = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	terminals = &sync.Map{}
)

func getTerm(id string) *os.File {
	d, ok := terminals.Load(id)
	if !ok {
		return nil
	}

	ptmx, ok := d.(*os.File)
	if !ok {
		return nil
	}

	return ptmx
}

func getTermSize(ctx *gin.Context) (uint16, uint16) {
	cols, _ := strconv.Atoi(ctx.Query("cols"))
	rows, _ := strconv.Atoi(ctx.Query("rows"))
	if cols < 1 || rows < 1 {
		cols = 80
		rows = 30
	}

	return uint16(cols), uint16(rows)
}

// handleShellList
// GET /api/v1/shell
func handleShellList(ctx *gin.Context) {
	shells := make([]string, 0)
	terminals.Range(func(key, value interface{}) bool {
		strId, _ := key.(string)
		shells = append(shells, strId)
		return true
	})

	utils.RespOkJSON(ctx, shells)
}

// handleShellCreate create a pty for web shell
// POST /api/v1/shell?cols={}&rows={}&shell={}
func handleShellCreate(ctx *gin.Context) {
	cols, rows := getTermSize(ctx)
	// Create arbitrary command.
	shell := ctx.Query("shell")
	if shell == "" {
		switch runtime.GOOS {
		case "windows":
			shell = "cmd.exe"
		default:
			shell = defaultShell
		}
	}
	c := exec.Command(shell)

	// Start the command with a pty.
	ptmx, err := pty.Start(c)
	if err != nil {
		log.E("start terminal failed", zap.Error(err))
		utils.RespJSON(ctx, -1, err.Error())
		return
	}
	id := strconv.Itoa(int(ptmx.Fd()))
	terminals.Store(id, ptmx)

	ws := &pty.Winsize{Cols: uint16(cols), Rows: uint16(rows)}
	pty.Setsize(ptmx, ws)

	utils.RespOkJSON(ctx, map[string]string{"id": id})
}

// handleShellSize resize pty
// POST /api/v1/shell/:id/size?cols={cols}&rows={rows}
func handleShellSize(ctx *gin.Context) {
	ptmx := getTerm(ctx.Param("id"))
	if ptmx == nil {
		utils.RespJSON(ctx, 1, "no such shell")
		return
	}

	cols, rows := getTermSize(ctx)
	pty.Setsize(ptmx, &pty.Winsize{
		Cols: uint16(cols), Rows: uint16(rows),
	})
}

// handleShellClose close shell
// DELETE /api/v1/shell/:id
func handleShellClose(ctx *gin.Context) {
	ptmx := getTerm(ctx.Param("id"))
	if ptmx == nil {
		utils.RespJSON(ctx, 1, "no such shell")
		return
	}
	terminals.Delete(ctx.Param("id"))
	utils.RespOkJSON(ctx)
}

// handleShellSession handle requests for web shell, make it a WebSocket
// GET /api/v1/terminal/:id
func handleShellSession(ctx *gin.Context) {
	ptmx := getTerm(ctx.Param("id"))
	if ptmx == nil {
		utils.RespErrJSON(ctx, http.StatusNotFound, 1, "no such terminal")
		return
	}

	c, err := up.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		utils.RespErrJSON(ctx, http.StatusBadRequest, 1, "open WebSocket failed")
		log.E("open shell WebSocket failed", zap.String("r_addr", ctx.Request.RemoteAddr), zap.Error(err))
		return
	}

	// close WebSocket while keep pty alive
	defer c.Close()

	rC, wC := make(chan []byte), make(chan []byte)

	// read from pty
	go func() {
		defer close(rC)
		s := bufio.NewScanner(ptmx)
		s.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			if atEOF && len(data) == 0 {
				return 0, nil, nil
			}
			return len(data), data[:], nil
		})

		for s.Scan() {
			rC <- s.Bytes()
		}
	}()

	// read from webSocket
	go func() {
		defer close(wC)
		for {
			_, data, err := c.ReadMessage()
			if err != nil {
				log.I("read from WebSocket failed", zap.String("r_addr", c.RemoteAddr().String()), zap.Error(err))
				return
			}
			wC <- data
		}
	}()

	for {
		select {
		case data, more := <-rC:
			if !more {
				return
			}

			// strip invalid utf-8 characters
			if !utf8.Valid(data) {
				s := string(data)
				v := make([]rune, 0, len(s))
				for i, r := range s {
					if r == utf8.RuneError {
						_, size := utf8.DecodeRuneInString(s[i:])
						if size == 1 {
							continue
						}
					}
					v = append(v, r)
				}
				data = []byte(string(v))
			}

			err := c.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.I("write to shell WebSocket failed", zap.String("r_addr", c.RemoteAddr().String()), zap.Error(err))
				return
			}
		case data, more := <-wC:
			if !more {
				return
			}

			_, err := ptmx.Write(data)
			if err != nil {
				log.I("write to shell terminal failed", zap.Error(err))
				return
			}
		}
	}
}
