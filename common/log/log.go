package log

import "github.com/rs/zerolog/log"

type Logger struct {
}

func Log(message string) {
	log.Printf("%v", message)
}
