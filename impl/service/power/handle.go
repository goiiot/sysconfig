package power

import (
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goiiot/sysconfig/cmd/server/conf"
	"github.com/goiiot/sysconfig/impl/service/utils"
)

func InitServicePower(v1 *gin.RouterGroup, config *conf.ServicePower) {
	if config.EnableReboot {
		v1.POST("/power/reboot", handleReboot)
	}

	if config.EnableShutdown {
		v1.POST("/power/shutdown", handleShutdown)
	}
}

var (
	rebootLock   int32
	shutdownLock int32
)

func getWaitTime(t string) (time.Duration, error) {
	var wait time.Duration
	if t == "" {
		wait = 0
		return wait, nil
	}

	ti, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return 0, err
	}

	return time.Until(ti), nil
}

// POST /api/v1/power/reboot?time={}
// time should be in RFC3339
func handleReboot(ctx *gin.Context) {
	t, err := getWaitTime(ctx.Query("time"))
	if err != nil {
		utils.RespErrJSON(ctx, http.StatusBadRequest, 1, "time format should be RFC3339")
		return
	}

	if !atomic.CompareAndSwapInt32(&rebootLock, 0, 1) {
		utils.RespErrJSON(ctx, http.StatusProcessing, 1, "system reboot is ongoing")
		return
	}

	errCh := reboot(t)
	go func() {
		timer := time.NewTimer(t)
		defer func() {
			atomic.StoreInt32(&rebootLock, 0)
			timer.Stop()
		}()

		select {
		case <-errCh:
			// TODO tell client if error happened
			return
		case <-timer.C:
			return
		}
	}()
}

// POST /api/v1/power/shutdown?time={}
func handleShutdown(ctx *gin.Context) {
	t, err := getWaitTime(ctx.Query("time"))
	if err != nil {
		utils.RespErrJSON(ctx, http.StatusBadRequest, 1, "time format should be RFC3339")
		return
	}

	if !atomic.CompareAndSwapInt32(&shutdownLock, 0, 1) {
		utils.RespErrJSON(ctx, http.StatusProcessing, 1, "system shutdown is ongoing")
		return
	}

	errCh := shutdown(t)
	go func() {
		timer := time.NewTimer(t)
		defer func() {
			atomic.StoreInt32(&shutdownLock, 0)
			timer.Stop()
		}()

		select {
		case <-errCh:
			// TODO tell client if error happened
			return
		case <-timer.C:
			return
		}
	}()
}
