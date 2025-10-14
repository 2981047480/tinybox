package log

import (
	"github.com/rs/zerolog/log"
)

type Logger struct {
}

func Log(message string) {
	log.Printf("%v", message)
}

func Warning(message string) {
	log.Warn().Msg(message)
}

func Error(message string) {
	log.Error().Msg(message)
}
