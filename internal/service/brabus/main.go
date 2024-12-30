package brabus

import (
	config "brabus/configs"
	"brabus/internal/yaml"
	"brabus/pkg/metrics"
	"github.com/mailru/easyjson"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func Run() {
	var rawBytes []byte
	var failCount int
	var globalConf config.Global

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to NATS")

	globalConf = yaml.UnmarshalGlobalConfig()
	failCount = 0

	for {
		time.Sleep(2 * time.Second)

		rawBytes, err = easyjson.Marshal(metrics.CollectMetrics())

		if err != nil {
			failCount++
			if failCount > globalConf.Limits.FailLimit {
				log.Fatalf("Fatal error marshalling metrics: %v", err)
			}

			log.Printf("Error during marshalling metrics: %v\n", err)
			continue
		}

		err = nc.Publish("metrics", rawBytes)
		if err != nil {
			failCount++
			if failCount > globalConf.Limits.FailLimit {
				log.Fatalf("Fatal error during publish metrics %v", err)
			}

			log.Printf("Error during publish metrics: %v\n", err)
		}
	}
}
