package banana

import (
	"brabus/internal/DTO"
	"brabus/pkg/analyze"
	"github.com/mailru/easyjson"
	"github.com/nats-io/nats.go"
	"log"
)

func Run() {
	metrics := new(DTO.Metrics)

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to NATS")

	_, err = nc.Subscribe("metrics", func(msg *nats.Msg) {
		err = easyjson.Unmarshal(msg.Data, metrics)
		if err != nil {
			log.Fatal(err)
		}

		analyze.ProcessMetrics(metrics)
	})

	if err != nil {
		log.Fatal(err)
	}

	select {}
}
