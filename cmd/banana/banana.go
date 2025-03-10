package main

import (
	"brabus/pkg/analyze"
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

	metrics := dto.Metrics{}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal().Stack().Err(logger.WrapError(err)).
			Str("NATS_URL", nats.DefaultURL).
			Msg("Error connecting to nats server")
	}
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
