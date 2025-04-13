package main

import (
	"brabus/pkg/app/brabus"
	"brabus/pkg/dto"
	"brabus/pkg/env"
	"brabus/pkg/logger"
	"brabus/pkg/yaml"
	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"log"
	"os"
	"time"
)

func main() {
	env.Init()
	zerolog, closeLog := logger.Init()
	defer closeLog()

	failCount := 0
	tick := time.Tick(2 * time.Second)
	quit := make(chan int)
	globalConf := yaml.UnmarshalGlobalConfig()

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		zerolog.Fatal().Stack().
			Str("NATS_URL", nats.DefaultURL).
			Err(errors.Wrap(err, "error connecting to nats server")).
			Send()
	}

	ctx := &dto.BrabusContext{
		FailCount:  &failCount,
		GlobalConf: globalConf,
		NC:         nc,
		Quit:       quit,
	}

	log.Println("Connected to NATS")

	for {
		select {
		case <-tick:
			brabus.ProcessMetrics(ctx, zerolog)
		case code := <-quit:
			log.Println("Quitting")
			nc.Close()
			os.Exit(code)
		}
	}
}
