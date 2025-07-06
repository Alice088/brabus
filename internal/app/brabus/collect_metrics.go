package brabus

import (
	"brabus/pkg/dto"
	"brabus/pkg/metrics"
	"github.com/mailru/easyjson"
	"github.com/rs/zerolog"
	"os"
	"time"
)

var (
	collectedMetrics *dto.Metrics
	err              error
	rawBytes         []byte
)

func (brabus *Brabus) CollectMetrics(logger zerolog.Logger, signal chan os.Signal) {
	tS := time.Now()
	collectedMetrics, err = metrics.Collect(logger)
	tE := time.Since(tS)

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Error collecting metrics")
		signal <- os.Kill
		return
	}

	logger.Debug().
		Str("Duration", tE.String()).
		Msgf("%+v", collectedMetrics)

	rawBytes, err = easyjson.Marshal(collectedMetrics)

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Failed to marshal metrics")
		signal <- os.Kill
		return
	}

	err = brabus.Nats.Publish("metrics", rawBytes)

	if err != nil {
		logger.Error().
			Err(err).
			Msg("Cannot publish metrics")
		signal <- os.Kill
		return
	}
}
