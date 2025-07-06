package main

import (
	"brabus/internal/app/brabus"
	"brabus/pkg/config"
	"brabus/pkg/log"
	"brabus/pkg/utils"
	"context"
	"github.com/avast/retry-go"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg = sync.WaitGroup{}

func main() {
	logger, closeLog := log.Init()
	defer closeLog()

	conf := config.Init()

	if conf.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		logger.Debug().Msgf("Config: %+v", conf)
	}

	utils.DefaultRetryConfig = utils.DefaultRetry(conf.Limit.Fail, 1, logger)

	defer func() {
		if err := recover(); err != nil {
			logger.Error().Msgf("Recovery: %v", err)
			wg.Wait()
			os.Exit(1)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	var connect *nats.Conn
	var err error
	err = retry.Do(func() error {
		connect, err = nats.Connect(nats.DefaultURL)
		return err
	}, utils.DefaultRetryConfig...)
	defer connect.Close()

	if err != nil {
		logger.Fatal().
			Str("NATS_URL", nats.DefaultURL).
			Err(err).
			Msg("error connecting to nats server")
	}

	logger.Info().Msg("Brabus started")

	signals := make(chan os.Signal, 1)
	defer close(signals)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	b := brabus.New(connect, conf)
	go b.Run(logger, &wg, brabus.Shutdown{
		Os:  signals,
		Ctx: ctx,
	})

	select {
	case <-signals:
		logger.Info().Msg("Stopping...")
		os.Exit(0)
	case <-ctx.Done():
		logger.Info().Msg("Stopping...")
		wg.Wait()
		os.Exit(0)
	}
}
