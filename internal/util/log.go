package util

import "github.com/rs/zerolog/log"

func LogError(msg string, err error) {
	log.Error().Msg(msg)
	log.Error().Err(err)
}

func LogDebug(msg string) {
	log.Debug().Msg(msg)
}