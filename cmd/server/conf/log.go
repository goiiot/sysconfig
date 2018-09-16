package conf

const (
	cfgLogLevel = "log.level"
)

var logFlags = []Flag{
	stringFlag(&config.Log.Level, cfgLogLevel, "l", "debug", "server log level"),
}

// LogConfig match log config block in yaml
type LogConfig struct {
	Level string
}
