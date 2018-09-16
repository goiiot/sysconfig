package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"github.com/goiiot/sysconfig/cmd/server/conf"
	"github.com/goiiot/sysconfig/impl/log"
	"github.com/goiiot/sysconfig/impl/service"
)

var (
	appName   = "i.MX6 Config Service"
	appDesc   = "Config i.MX6 with ease"
)

var (
	// fields below will be replaced while built with goreleaser
	// according to .goreleaser.yaml
	version   = "snapshot"
	commit    = "none"
	buildTime = "none"
	goVersion = "none"
)

var app = &cobra.Command{
	Use:     appName,
	Short:   appDesc,
	Version: version,
	Run:     start,
}

func main() {
	conf.SetFlags(app)
	// execute will parse flags and run the `start` function below
	if err := app.Execute(); err != nil {
		log.I("server exit", zap.Error(err))
	}
}

func start(app *cobra.Command, args []string) {
	conf.ParseConfig(app, args)

	cfg := conf.GetConfig()
	cfg.Version = version
	cfg.Commit = commit
	cfg.BuildTime = buildTime
	cfg.GoVersion = goVersion

	ctx := context.WithValue(context.Background(), conf.KeyConfig, cfg)
	appCtx, exit := context.WithCancel(ctx)

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, os.Kill)
	go func() {
		s := <-sigCh
		log.I("user interrupted", zap.String("signal", s.String()))
		exit()
	}()

	service.Start(appCtx)
}
