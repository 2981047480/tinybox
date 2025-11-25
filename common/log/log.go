package log

import (
	"github.com/rs/zerolog/log"
)

type logger struct {
}

func NewLogger() logger {
	return logger{}
}

var l logger = NewLogger()

func (l *logger) Log(message string) {
	log.Printf("%v", message)
}

func (l *logger) Warning(message string) {
	log.Warn().Msg(message)
}

func (l *logger) Error(message string) {
	log.Error().Msg(message)
}

// 包级函数，供外部直接调用
func Log(message string) {
	l.Log(message)
}

func Warning(message string) {
	l.Warning(message)
}

func Error(message string) {
	l.Error(message)
}
