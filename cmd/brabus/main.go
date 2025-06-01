package main

import (
	app "brabus/internal/app/brabus"
	"brabus/pkg/env"
	"brabus/pkg/exretry"
	"brabus/pkg/log"
	"brabus/pkg/yaml"
	"context"
	"github.com/avast/retry-go"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
)

var (
	connect *nats.Conn
	err     error
)

func main() {
	env.Init()

	logger, closeLog := log.Init()
	defer closeLog()

	config := yaml.UnmarshalGlobalConfig()

	if config.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	exretry.DefaultRetryConfig = exretry.DefaultRetry(config.Limits.FailLimit, 1, logger)

	defer func() {
		if err := recover(); err != nil {
			logger.Error().Msgf("Recovery: %v", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	logger.Info().Msg("connecting to NATS...")
	err = retry.Do(func() error {
		connect, err = nats.Connect(nats.DefaultURL)
		return err
	}, exretry.DefaultRetryConfig...)

	if err != nil {
		logger.Fatal().
			Str("NATS_URL", nats.DefaultURL).
			Err(err).
			Msg("error connecting to nats server")
	}

	logger.Info().Msg("Starting Brabus....")
	brabus := app.NewBrabus(ctx, stop, connect, config)

	brabus.Scan(logger)

	select {
	case <-ctx.Done():
		logger.Info().Msg("Stoping program.")
		connect.Close()
		os.Exit(0)
	}
}
