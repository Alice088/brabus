package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"os"

	"time"
)

type CloseFunc func()

func Init() (*zerolog.Logger, CloseFunc) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	logFile, err := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}
	multi := zerolog.MultiLevelWriter(consoleWriter, logFile)
	logger := zerolog.New(multi).With().Timestamp().Logger()
	logger.Debug().Msg("Logger initialized")

	return &logger, func() {
		Close(logFile)
	}
}
