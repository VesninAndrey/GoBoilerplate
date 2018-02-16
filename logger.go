package main

import (
	"os"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

var log zerolog.Logger

func initLogger()  {
	log = zlog.Logger.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	log.Info().Int("Build", version).Msg("Hello!")
}
