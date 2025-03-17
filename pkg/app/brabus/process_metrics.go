package brabus

import (
	"brabus/pkg/dto"
	"brabus/pkg/metrics"
	"github.com/mailru/easyjson"
	"github.com/pkg/errors"
	zl "github.com/rs/zerolog"
)

func ProcessMetrics(ctx *dto.BrabusContext, zerolog *zl.Logger) {
	rawBytes, err := easyjson.Marshal(metrics.Collect())

	if err != nil {
		*ctx.FailCount++
		if *ctx.FailCount > ctx.GlobalConf.Limits.FailLimit {
			zerolog.Error().Stack().
				Int("TRY", *ctx.FailCount).
				Err(errors.Wrap(err, "limit of errors! Marshal error")).
				Send()
			ctx.Quit <- 1
		}

		zerolog.Warn().Stack().
			Int("TRY", *ctx.FailCount).
			Err(errors.Wrap(err, "error during marshalling metrics. Marshal error")).
			Send()
	}

	err = ctx.NC.Publish("metrics", rawBytes)
	if err != nil {
		*ctx.FailCount++
		if *ctx.FailCount > ctx.GlobalConf.Limits.FailLimit {
			zerolog.Fatal().Stack().
				Int("TRY", *ctx.FailCount).
				Err(errors.Wrap(err, "Fatal error during publish metrics. Publish error")).
				Send()

			ctx.Quit <- 1
		}

		zerolog.Warn().Stack().
			Int("TRY", *ctx.FailCount).
			Err(errors.Wrap(err, "Error during publish metrics. Publish error")).
			Send()
	}
}
