package conf

import (
	"io/ioutil"
	"strings"

	"github.com/goiiot/sysconfig/impl/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

const (
	KeyConfig = "config"
)

var (
	confFile string
	config   = new(Config)
)

// Flags return all flags used by this app
func SetFlags(cmd *cobra.Command) {
	flags := []Flag{
		stringFlag(&confFile, "config", "c", "config.yaml", "configuration file"),
	}
	flags = append(flags, serverFlags...)
	flags = append(flags, serviceFlags...)
	flags = append(flags, logFlags...)

	fs := cmd.Flags()
	for _, f := range flags {
		f(fs)
	}
}

func ParseConfig(c *cobra.Command, args []string) {
	// init as fatal at first, then switch according to user configuration
	log.Init(log.Fatal)
	logLevel := func() log.Level {
		switch strings.ToLower(strings.TrimSpace(config.Log.Level)) {
		case "debug":
			return log.Debug
		case "info":
			return log.Info
		case "warning":
			return log.Warning
		case "error":
			return log.Error
		case "fatal":
			return log.Fatal
		case "silent":
			return log.Silent
		default:
			return log.Debug
		}
	}()

	f, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.F("failed to read config", false, zap.Error(err))
	}

	if err = yaml.Unmarshal(f, config); err != nil {
		log.F("parse config file failed", false, zap.Error(err))
	}

	// init user capability
	if users := config.Server.Auth.Users; users != nil && len(users) > 0 {
		for i := range users {
			users[i].initCapability()
		}
	}

	if err = c.ParseFlags(args); err != nil {
		log.F("invalid args", false, zap.Error(err))
	}

	log.Init(logLevel)
}

type Config struct {
	Server  ServerConfig
	Service ServiceConfig
	Log     LogConfig
	// app version
	Version   string
	Commit    string
	BuildTime string
	GoVersion string
}

func GetConfig() *Config {
	return config
}
