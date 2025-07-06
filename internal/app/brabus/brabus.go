package brabus

import (
	"brabus/pkg/config"
	"github.com/nats-io/nats.go"
)

type Brabus struct {
	Nats   *nats.Conn
	Config config.Config
}

func New(nats *nats.Conn, config config.Config) Brabus {
	return Brabus{
		Nats:   nats,
		Config: config,
	}
}
