package brabus

import (
	"brabus/pkg/metrics"
	"github.com/mailru/easyjson"
	"github.com/rs/zerolog"
)

func (brabus *Brabus) ProcessMetrics(logger *zerolog.Logger) {
	collectedMetrics, err := metrics.Collect()
	if err != nil {
		logger.Error().Err(err).Msg("Error collecting metrics")
		brabus.stop()
		return
	}

	logger.Info().Msgf("%+v", collectedMetrics)

	rawBytes, err := easyjson.Marshal(collectedMetrics)
	if err != nil {
		logger.Warn().
			Err(err).
			Msg("Failed to marshal metrics")
		brabus.stop()
		return
	}

	err = brabus.Nats.Publish("metrics", rawBytes)
	if err != nil {
		logger.Warn().
			Err(err).
			Msg("Cannot publish metrics")
		brabus.stop()
		return
	}
}
