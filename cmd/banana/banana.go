package main

import (
	"brabus/internal/app/banana/analyze"
	"brabus/internal/app/banana/storage/badger"
	"brabus/pkg/config"
	"brabus/pkg/dto"
	"brabus/pkg/env"
	"brabus/pkg/log"
	"brabus/pkg/utils"
	"github.com/avast/retry-go"
	bad "github.com/dgraph-io/badger/v4"
	"github.com/mailru/easyjson"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

func main() {
	env.Init()
	logger, closeLog := log.Init()
	defer closeLog()

	conf := config.Init()

	if conf.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		logger.Debug().Msgf("Config: %+v", conf)
	}

	db, err := badger.Init()

	if err != nil {
		logger.Fatal().Stack().Err(err).Msg("Error init badger")
	}

	defer func(db *bad.DB) {
		err := db.Close()
		if err != nil {
			logger.Error().Err(err).Msg("Error closing db")
		}
	}(db)

	metrics := dto.Metrics{}

	utils.DefaultRetryConfig = utils.DefaultRetry(conf.Limit.Fail, 1, logger)

	var connect *nats.Conn
	err = retry.Do(func() error {
		connect, err = nats.Connect(nats.DefaultURL)
		return err
	}, utils.DefaultRetryConfig...)

	if err != nil {
		logger.Fatal().
			Str("NATS_URL", nats.DefaultURL).
			Err(err).
			Msg("error connecting to nats server")
	}

	defer connect.Close()

	logger.Info().Msg("Connected to NATS")

	_, err = nc.Subscribe("metrics", func(msg *nats.Msg) {
		err = easyjson.Unmarshal(msg.Data, &metrics)
		if err != nil {
			logger.Fatal().Stack().Err(logger.WrapError(err)).
				Str("NATS_URL", nats.DefaultURL).
				Msg("Error unmarshalling metrics")
		}

		analyze.Metrics(metrics)
	})

	if err != nil {
		logger.Fatal().Stack().Err(logger.WrapError(err)).
			Str("NATS_URL", nats.DefaultURL).
			Msg("Error subscribing to NATS")
	}

	select {}
}
