package service

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goiiot/sysconfig/cmd/sysconfig/conf"
	"github.com/goiiot/sysconfig/impl/log"
	"github.com/rakyll/statik/fs"
	"go.uber.org/zap"

	_ "github.com/goiiot/sysconfig/impl/ui"
)

func startWebServer(appCtx context.Context, r *gin.Engine, cfg *conf.ServerWebConfig) {
	if cfg.Root == "bundled" {
		// use bundled web app
		appFS, err := fs.New()
		if err != nil {
			log.F("bundled web app selected, but no bundled file found", false, zap.Error(err))
		}
		r.StaticFS("/app", appFS)
	} else {
		// use external web app
		r.Static("/app", cfg.Root)
	}

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/app")
	})

	// start http server if enabled
	var httpSrv *http.Server
	if cfg.HTTP.Enabled {
		httpSrv = &http.Server{
			Addr:    cfg.HTTP.Listen,
			Handler: r,
		}
		go func() {
			if err := httpSrv.ListenAndServe(); err != nil {
				log.I("http server exit", zap.Error(err))
			}
		}()
	}

	// start https server if enabled
	var httpsSrv *http.Server
	if cfg.HTTPS.Enabled {
		httpsSrv = &http.Server{
			Addr:    cfg.HTTPS.Listen,
			Handler: r,
		}
		go func() {
			if err := httpsSrv.ListenAndServeTLS(cfg.HTTPS.TLSCert, cfg.HTTPS.TLSKey); err != nil {
				log.I("https server exit", zap.Error(err))
			}
		}()
	}

	<-appCtx.Done()
	shutdownCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if httpSrv != nil {
		if err := httpSrv.Shutdown(shutdownCtx); err != nil {
			log.I("http server shutdown err", zap.Error(err))
		}
	}

	if httpsSrv != nil {
		if err := httpsSrv.Shutdown(shutdownCtx); err != nil {
			log.I("https server shutdown err", zap.Error(err))
		}
	}
}
