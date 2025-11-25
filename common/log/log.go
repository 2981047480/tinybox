package log

import (
	"github.com/rs/zerolog/log"
)

type Logger struct {
}

func NewLogger() Logger {
	return Logger{}
}

var L Logger = NewLogger()

func (l *Logger) Log(message string) {
	log.Printf("%v", message)
}

func (l *Logger) Warning(message string) {
	log.Warn().Msg(message)
}

func (l *Logger) Error(message string) {
	log.Error().Msg(message)
}
