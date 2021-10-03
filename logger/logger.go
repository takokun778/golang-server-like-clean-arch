package logger

import (
	"os"
	"strconv"

	"go.uber.org/zap"
)

type logger struct {
	*zap.Logger
}

// シングルトン
var singleton = newLogger()

func newLogger() *logger {
	env := os.Getenv("ENV")
	if env == "prod" {
		prod, _ := zap.NewProduction()
		return &logger{prod}
	}
	zap, _ := zap.NewDevelopment()
	return &logger{zap}
}

func toFields(v ...interface{}) []zap.Field {
	fields := make([]zap.Field, 0)
	for i, f := range v {
		fields = append(fields, zap.Any(strconv.Itoa(i), f))
	}
	return fields
}

// Debugレベル
func Debug(msg string, v ...interface{}) {
	env := os.Getenv("ENV")
	if env == "prod" {
		return
	}
	if v != nil {
		singleton.Debug(msg, toFields(v)...)
		return
	}
	singleton.Debug(msg)
}

// Infoレベル
func Info(msg string, v ...interface{}) {
	if v != nil {
		singleton.Info(msg, toFields(v)...)
		return
	}
	singleton.Info(msg)
}

// Warnレベル
func Warn(msg string, v ...interface{}) {
	if v != nil {
		singleton.Warn(msg, toFields(v)...)
		return
	}
	singleton.Warn(msg)
}

// Errorレベル
func Error(msg string, v ...interface{}) {
	if v != nil {
		singleton.Error(msg, toFields(v)...)
		return
	}
	singleton.Error(msg)
}

func StartLog() {
	singleton.Info("start server")
}
