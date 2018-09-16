package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Ok     bool        `json:"ok"`
	Code   int         `json:"code"`
	Reason string      `json:"reason,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func RespErrJSON(ctx *gin.Context, httpCode int, code int, reason string) {
	ctx.JSON(httpCode, &response{
		Code:   code,
		Ok:     false,
		Reason: reason,
	})
}

func RespOkJSON(ctx *gin.Context, data ...interface{}) {
	ctx.JSON(http.StatusOK, &response{
		Code: 0,
		Ok:   true,
		Data: func() interface{} {
			if n := len(data); n == 0 {
				return nil
			} else if n == 1 {
				return data[0]
			} else {
				return data
			}
		}(),
	})
}

func RespJSON(ctx *gin.Context, code int, reason string, data ...interface{}) {
	ctx.JSON(http.StatusOK, &response{
		Code:   code,
		Ok:     code == 0,
		Reason: reason,
		Data: func() interface{} {
			if n := len(data); n == 0 {
				return nil
			} else if n == 1 {
				return data[0]
			} else {
				return data
			}
		}(),
	})
}
