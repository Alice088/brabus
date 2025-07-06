package main

import (
	"brabus/internal/app/banana/analyze"
	"brabus/internal/app/banana/storage/badger"
	"brabus/pkg/config"
	"brabus/pkg/dto"
	"brabus/pkg/log"
	"brabus/pkg/utils"
	"fmt"
	"github.com/avast/retry-go"
	bad "github.com/dgraph-io/badger/v4"
	"github.com/mailru/easyjson"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

func main() {
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

	metrics := make(chan *nats.Msg)
	_, err = connect.ChanSubscribe("metrics", metrics)
	if err != nil {
		logger.Fatal().
			Err(err).
			Msg("Error subscribing to NATS")
	}

	fmt.Println("--------------------------------WARNINGS--------------------------------------")

	for {
		select {
		case msg := <-metrics:
			var metric dto.Metrics
			err = easyjson.Unmarshal(msg.Data, &metric)
			if err != nil {
				logger.Error().
					Err(err).
					Msg("Error unmarshalling metrics")
			}

			logger.Debug().Msgf("Metrics received: %+v", metric)

			analyze.Metrics(metric, logger)
		}
	}
}
