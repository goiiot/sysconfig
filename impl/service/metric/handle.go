package metric

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"github.com/goiiot/sysconfig/cmd/sysconfig/conf"
	"github.com/goiiot/sysconfig/impl/log"
	"github.com/goiiot/sysconfig/impl/service/utils"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var (
	refreshInterval time.Duration
	up              = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
)

func InitServiceMetric(v1 *gin.RouterGroup, config *conf.ServiceMonitoring) {
	if !config.Enabled {
		return
	}

	refreshInterval = config.RefreshInterval

	v1.GET("/metrics/aggregated/all", handleGetAllMetrics)
	v1.GET("/metrics/aggregated/static", handleGetStaticMetrics)
	v1.GET("/metrics/aggregated/dynamic", handleGetDynamicMetrics)

	v1.GET("/metrics/cpu", handleGetCpuMetrics)
	v1.GET("/metrics/cpu/info", handleGetCpuInfoMetrics)
	v1.GET("/metrics/cpu/usage", handleGetCpuUsageMetrics)

	v1.GET("/metrics/mem", handleGetMemMetrics)

	v1.GET("/metrics/disk", handleGetDiskMetrics)
	v1.GET("/metrics/disk/info", handleGetDiskInfoMetrics)
	v1.GET("/metrics/disk/usage", handleGetDiskUsageMetrics)

	v1.GET("/metrics/net", handleGetNetMetrics)
	v1.GET("/metrics/net/info", handleGetNetInterfaceMetrics)
	v1.GET("/metrics/net/usage", handleGetNetUsageMetrics)

	v1.GET("/metrics/host", handleGetHostMetrics)

	v1.GET("/metrics/process", handleGetProcessMetrics)
}

func handleGetAllMetrics(ctx *gin.Context) {
	m := &Metrics{}
	setAllMetrics(m)
	utils.RespOkJSON(ctx, m)
}

func handleGetStaticMetrics(ctx *gin.Context) {
	m := &Metrics{}
	setAllStaticMetrics(m)
	utils.RespOkJSON(ctx, m)
}

func handleGetDynamicMetrics(ctx *gin.Context) {
	m := &Metrics{}
	setAllDynamicMetrics(m)
	utils.RespOkJSON(ctx, m)
}

func handleGetCpuMetrics(ctx *gin.Context) {
	m := &Metrics{}
	setCpuInfoMetrics(m)
	setCpuUsageMetrics(m)
	utils.RespOkJSON(ctx, m)
}

func handleGetCpuInfoMetrics(ctx *gin.Context) {
	m := &Metrics{}
	setCpuInfoMetrics(m)
	utils.RespOkJSON(ctx, m)
}

func handleGetCpuUsageMetrics(ctx *gin.Context) {
	m := &Metrics{}
	setCpuUsageMetrics(m)
	utils.RespOkJSON(ctx, m)
}

func handleGetMemMetrics(ctx *gin.Context) {
	m := &Metrics{}
	setMemMetrics(m)
	utils.RespOkJSON(ctx, m)
}

func handleGetDiskMetrics(ctx *gin.Context) {
	m := &Metrics{}
	setDiskInfoMetrics(m)
	setDiskUsageMetrics(m)
	utils.RespOkJSON(ctx, m)
}

func handleGetDiskInfoMetrics(ctx *gin.Context) {
	m := &Metrics{}
	setDiskInfoMetrics(m)
	utils.RespOkJSON(ctx, m)
}

func handleGetDiskUsageMetrics(ctx *gin.Context) {
	m := &Metrics{}
	setDiskUsageMetrics(m)
	utils.RespOkJSON(ctx, m)
}

func handleGetHostMetrics(ctx *gin.Context) {
	m := &Metrics{}
	setHostMetrics(m)
	utils.RespOkJSON(ctx, m)
}

func handleGetNetMetrics(ctx *gin.Context) {
	m := &Metrics{}
	setNetInterfaceMetrics(m)
	setNetUsageMetrics(m)
	utils.RespOkJSON(ctx, m)
}

func handleGetNetInterfaceMetrics(ctx *gin.Context) {
	m := &Metrics{}
	setNetInterfaceMetrics(m)
	utils.RespOkJSON(ctx, m)
}

func handleGetNetUsageMetrics(ctx *gin.Context) {
	m := &Metrics{}
	setNetUsageMetrics(m)
	utils.RespOkJSON(ctx, m)
}

func handleGetProcessMetrics(ctx *gin.Context) {
	m := &Metrics{}
	setProcessMetrics(m)
	utils.RespOkJSON(ctx, m)
}

func handleMetricSession(ctx *gin.Context) {
	c, err := up.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		utils.RespErrJSON(ctx, http.StatusBadRequest, 1, "open WebSocket failed")
		log.E("open metric WebSocket failed", zap.String("r_addr", ctx.Request.RemoteAddr), zap.Error(err))
		return
	}
	defer c.Close()

	t := time.NewTicker(refreshInterval)
	defer t.Stop()

	ch := make(chan *Metrics)
	go func() {
		for range t.C {
			ch <- getCurrentMetrics()
		}
	}()

	for metrics := range ch {
		m, err := json.Marshal(metrics)
		if err != nil {
			log.E("marshal metrics json failed", zap.Error(err))
			continue
		}

		if err := c.WriteMessage(websocket.TextMessage, m); err != nil {
			log.I("write to metrics WebSocket failed", zap.String("r_addr", c.RemoteAddr().String()), zap.Error(err))
			return
		}
	}
}
