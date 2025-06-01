package main

import (
	app "brabus/internal/app/brabus"
	"brabus/pkg/env"
	"brabus/pkg/log"
	"context"
	"github.com/nats-io/nats.go"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	env.Init()

	logger, closeLog := log.Init()
	defer closeLog()

	defer func() {
		if err := recover(); err != nil {
			logger.Error().Msgf("Recovery: %v", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	connect, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		logger.Fatal().
			Str("NATS_URL", nats.DefaultURL).
			Err(err).
			Msg("error connecting to nats server")
	}

	logger.Info().Msg("Starting Brabus....")
	brabus := app.NewBrabus(ctx, stop, connect)

	brabus.Scan(logger)

	select {
	case <-ctx.Done():
		logger.Info().Msg("Stoping program.")
		connect.Close()
		os.Exit(0)
	}
}
