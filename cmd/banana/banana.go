package main

import (
	"brabus/pkg/app/banana/analyze"
	"brabus/pkg/app/banana/storage/badger"
	"brabus/pkg/dto"
	"brabus/pkg/env"
	"brabus/pkg/logger"
	"github.com/mailru/easyjson"
	"github.com/nats-io/nats.go"
)

func main() {
	env.Init()
	log, closeLog := logger.Init()
	defer closeLog()

	zerolog, closeLog := logger.Init()
	defer closeLog()

	db, err := badger.Init()

	if err != nil {
		zerolog.Fatal().Stack().Err(err).Msg("Error init badger")
	}

	defer db.Close()

	metrics := dto.Metrics{}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal().Stack().Err(logger.WrapError(err)).
			Str("NATS_URL", nats.DefaultURL).
			Msg("Error connecting to nats server")
	}

	defer nc.Close()

	log.Info().Msg("Connected to NATS")

	_, err = nc.Subscribe("metrics", func(msg *nats.Msg) {
		err = easyjson.Unmarshal(msg.Data, &metrics)
		if err != nil {
			log.Fatal().Stack().Err(logger.WrapError(err)).
				Str("NATS_URL", nats.DefaultURL).
				Msg("Error unmarshalling metrics")
		}

		analyze.Metrics(metrics)
	})

	if err != nil {
		log.Fatal().Stack().Err(logger.WrapError(err)).
			Str("NATS_URL", nats.DefaultURL).
			Msg("Error subscribing to NATS")
	}

	select {}
}
