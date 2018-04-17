package log

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

func (l *Logger) WithError(err error) *Logger {
	return &Logger{l.With("error", err)}
}

// func (l *Logger) With(msg string, value interface{}) {

// }
