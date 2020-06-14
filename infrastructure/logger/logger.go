package logger

import (
	"errors"
)

type Logger interface {
	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	Fatalln(args ...interface{})

	WithFields(keyValues Fields) Logger

	WithError(err error) Logger
}

type Fields map[string]interface{}

const (
	//Debug has verbose message
	Debug = "debug"
	//Info is default log level
	Info = "info"
	//Warn is for logging messages about possible issues
	Warn = "warn"
	//Error is for logging errors
	Error = "error"
	//Fatal is for logging fatal messages. The sytem shutsdown after logging the message.
	Fatal = "fatal"
)

const (
	InstanceZapLogger int = iota
	InstanceLogrusLogger
)

var (
	errInvalidLoggerInstance = errors.New("invalid logger instance")
)

//NewLogger retorna a instância de um logger
func NewLogger(level string, formatJSON bool, loggerInstance int) (Logger, error) {
	switch loggerInstance {
	case InstanceZapLogger:
		var logger = NewZapLogger(level, formatJSON)
		return logger, nil

	case InstanceLogrusLogger:
		var logger = NewLogrusLogger(level, formatJSON)
		return logger, nil

	default:
		return nil, errInvalidLoggerInstance
	}
}
