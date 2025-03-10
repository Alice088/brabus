package main

import (
	config "brabus/pkg/dto"
	"brabus/pkg/env"
	"brabus/pkg/logger"
	"brabus/pkg/metrics"
	"brabus/pkg/yaml"
	"github.com/mailru/easyjson"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"time"
)

func main() {
	env.Init()
	log, closeLog := logger.Init()
	defer closeLog()

	var rawBytes []byte
	var failCount int
	var globalConf config.Global

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal().Stack().Err(errors.Wrap(err, err.Error())).Str("NATS_URL", nats.DefaultURL).Msg("Error connecting to nats server")
	}
	log.Info().Msg("Connected to NATS")

	globalConf = yaml.UnmarshalGlobalConfig()
	failCount = 0

	for {
		time.Sleep(2 * time.Second)

		rawBytes, err = easyjson.Marshal(metrics.Collect())

		if err != nil {
			failCount++
			if failCount > globalConf.Limits.FailLimit {
				log.Fatal().Err(err).Msg("Fatal error marshalling metrics")
			}

			log.Error().Err(err).Msg("Error during marshalling metrics")
			continue
		}

		err = nc.Publish("metrics", rawBytes)
		if err != nil {
			failCount++
			if failCount > globalConf.Limits.FailLimit {
				log.Fatal().Err(err).Msg("Fatal error during publish metrics")
			}

			log.Error().Err(err).Msg("Error during publish metrics")
		}
	}
}
