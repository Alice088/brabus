package exretry

import (
	"github.com/avast/retry-go"
	"github.com/rs/zerolog"
	"time"
)

var (
	DefaultRetryConfig []retry.Option
)

func DefaultRetry(attempts uint8, seconds time.Duration, logger *zerolog.Logger) []retry.Option {
	return []retry.Option{
		retry.Attempts(uint(attempts)),
		retry.OnRetry(func(n uint, err error) {
			logger.Warn().Uint("TRY", n+1).Msg("Retrying...")
		}),
		retry.Delay(seconds * time.Second),
		retry.LastErrorOnly(true),
	}
}
