package service

import (
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/goiiot/sysconfig/cmd/sysconfig/conf"
	"github.com/goiiot/sysconfig/impl/service/auth"
	"github.com/goiiot/sysconfig/impl/service/configure"
	"github.com/goiiot/sysconfig/impl/service/file"
	"github.com/goiiot/sysconfig/impl/service/metric"
	"github.com/goiiot/sysconfig/impl/service/power"
	"github.com/goiiot/sysconfig/impl/service/shell"
	"github.com/goiiot/sysconfig/impl/service/utils"
)

// Start services according to configuration
func Start(appCtx context.Context) {
	config, _ := appCtx.Value(conf.KeyConfig).(*conf.Config)

	r := gin.Default()
	// allow all origins
	r.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders: []string{
			"Origin", "Content-Length", "Content-Type",
			"Access-Control-Allow-Origin", "Content-Disposition"},
	}))

	// handle 404
	r.NoRoute(func(c *gin.Context) { utils.RespJSON(c, http.StatusNotFound, "page not found") })

	// api v1
	v1 := r.Group("/api/v1")
	v1.GET("/cap", handleCapability(&config.Service))
	v1.GET("/version", handleVersion(config))

	// init auth first
	auth.InitMiddlewareAuth(v1, &config.Server.Auth)

	// init services according to related configuration
	metric.InitServiceMetric(v1, &config.Service.Monitoring)
	shell.InitServiceShell(v1, &config.Service.Shell)
	file.InitServiceFile(v1, &config.Service.File)
	configure.InitServiceConfig(v1, &config.Service.Configure)
	power.InitServicePower(v1, &config.Service.Power)

	// start web server according to config
	startWebServer(appCtx, r, &config.Server.Web)
}

func handleCapability(config *conf.ServiceConfig) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// TODO translate config `Enabled` fields to capability
	}
}

func handleVersion(config *conf.Config) func(*gin.Context) {
	return func(ctx *gin.Context) {
		utils.RespOkJSON(ctx, gin.H{
			"version":    config.Version,
			"commit":     config.Commit,
			"go_version": config.GoVersion,
			"build_time": config.BuildTime,
		})
	}
}
