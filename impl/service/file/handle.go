package file

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/goiiot/sysconfig/cmd/sysconfig/conf"
	"github.com/goiiot/sysconfig/impl/service/utils"
	"github.com/mholt/archiver"
)

// InitServiceFile init file upload/download service
func InitServiceFile(v1 *gin.RouterGroup, config *conf.ServiceFile) {
	if config.EnableDownload {
		v1.GET("/file", handleFileDownload)
	}

	if config.EnableUpload {
		v1.PUT("/file", handleFileUpload)
	}
}

// handleFileUpload upload file/directory to server via http(s)
// POST /api/v1/file?dst={}
func handleFileUpload(ctx *gin.Context) {
	dst := ctx.Query("dst")
	if len(strings.TrimSpace(dst)) == 0 {
		utils.RespErrJSON(ctx, http.StatusBadRequest, 1, "dst path is empty")
		return
	}

	if !filepath.IsAbs(dst) {
		utils.RespErrJSON(ctx, http.StatusBadRequest, 1, fmt.Sprintf("dst path %s is not absolute", dst))
		return
	}

	f, err := os.Stat(dst)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(dst, os.FileMode(0755))
			if err != nil {
				utils.RespErrJSON(ctx, http.StatusBadRequest, 1, err.Error())
				return
			}
			f, err = os.Stat(dst)
		} else {
			utils.RespErrJSON(ctx, http.StatusBadRequest, 1, err.Error())
			return
		}
	}

	if !f.IsDir() {
		utils.RespErrJSON(ctx, http.StatusBadRequest, 1, "dst path not a directory")
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		utils.RespErrJSON(ctx, http.StatusBadRequest, 1, err.Error())
		return
	}

	files := form.File["files"]
	for _, file := range files {
		if err := ctx.SaveUploadedFile(file, filepath.Join(dst, file.Filename)); err != nil {
			utils.RespErrJSON(ctx, http.StatusNotAcceptable, 1, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
	}

	utils.RespOkJSON(ctx)
}

// handleFileDownload download files (archived into a zip/tarball file) from server via http(s)
// GET /api/v1/file?src={}&format={}
func handleFileDownload(ctx *gin.Context) {
	src := ctx.Query("src")
	format := ctx.Query("format")
	ar, suffix, mime := func() (a archiver.Archiver, suffix, mime string) {
		switch strings.ToLower(strings.TrimSpace(format)) {
		case "zip":
			return archiver.Zip, "zip", "zip"
		case "tar":
			return archiver.Tar, "tar", "tar"
		case "tar.gz":
			return archiver.TarGz, "tar.gz", "tar"
		case "tar.bz2":
			return archiver.TarBz2, "tar.bz2", "tar"
		case "tar.xz":
			return archiver.TarXZ, "tar.xz", "tar"
		default:
			return archiver.Zip, "zip", "zip"
		}
	}()

	if !filepath.IsAbs(src) {
		utils.RespErrJSON(ctx, http.StatusBadRequest, 1, fmt.Sprintf("src path %s is not absolute", src))
		return
	}

	_, err := os.Stat(src)
	if err != nil && os.IsNotExist(err) {
		utils.RespErrJSON(ctx, http.StatusBadRequest, 1, fmt.Sprintf("src path %s doesn't exist", src))
		return
	}

	ctx.Header("Content-Type", fmt.Sprintf("application/%s", mime))
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.%s\"", filepath.Base(src), suffix))
	ar.Write(ctx.Writer, []string{src})
}
