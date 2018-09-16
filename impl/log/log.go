package log

import (
	"go.uber.org/zap"
)

type Level int

const (
	Debug Level = iota
	Info
	Warning
	Error
	Fatal
	Silent
)

var (
	level Level
	log   *zap.Logger
)

func Init(l Level) {
	level = l
	if level == Silent {
		return
	}

	if level == Debug {
		log, _ = zap.NewDevelopment(zap.AddCallerSkip(1))
	} else {
		log, _ = zap.NewProduction(zap.AddCallerSkip(1))
	}
}

// D debug
func D(msg string, fields ...zap.Field) {
	if level <= Debug && log != nil {
		log.Debug(msg, fields...)
	}
}

// I info
func I(msg string, fields ...zap.Field) {
	if level <= Info && log != nil {
		log.Info(msg, fields...)
	}
}

// W warning
func W(msg string, fields ...zap.Field) {
	if level <= Warning && log != nil {
		log.Warn(msg, fields...)
	}
}

// E error
func E(msg string, fields ...zap.Field) {
	if level <= Error && log != nil {
		log.Error(msg, fields...)
	}
}

// F fatal
func F(msg string, rec bool, fields ...zap.Field) (rst interface{}) {
	if level <= Fatal && log != nil {
		if rec {
			fields = append(fields, zap.Bool("recover", true))
			defer func() {
				rst = recover()
			}()
		}
		log.Panic(msg, fields...)
	}

	return nil
}
