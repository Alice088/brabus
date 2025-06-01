package brabus

import (
	"brabus/pkg/dto"
	"brabus/pkg/exretry"
	"brabus/pkg/metrics"
	"github.com/avast/retry-go"
	"github.com/mailru/easyjson"
	"github.com/rs/zerolog"
)

var (
	collectedMetrics *dto.Metrics
	err              error
	rawBytes         []byte
)

func (brabus *Brabus) ProcessMetrics(logger *zerolog.Logger) {
	logger.Debug().Msg("Collecting metrics...")
	err = retry.Do(func() error {
		collectedMetrics, err = metrics.Collect()
		return err
	}, exretry.DefaultRetryConfig...)

	if err != nil {
		logger.Error().Err(err).Msg("Error collecting metrics")
		brabus.stop()
		return
	}

	logger.Info().Msgf("%+v", collectedMetrics)

	err = retry.Do(func() error {
		rawBytes, err = easyjson.Marshal(collectedMetrics)
		return err
	}, exretry.DefaultRetryConfig...)

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("Failed to marshal metrics")
		brabus.stop()
		return
	}

	err = retry.Do(func() error {
		return brabus.Nats.Publish("metrics", rawBytes)
	}, exretry.DefaultRetryConfig...)

	if err != nil {
		logger.Warn().
			Err(err).
			Msg("Cannot publish metrics")
		brabus.stop()
		return
	}
}
