package brabus

import (
	"brabus/pkg/dto"
	"context"
	"github.com/nats-io/nats.go"
)

type Brabus struct {
	Nats   *nats.Conn
	Config dto.Config
	ctx    context.Context
	stop   context.CancelFunc
}

func NewBrabus(ctx context.Context, cancelFunc context.CancelFunc, nats *nats.Conn, config dto.Config) Brabus {
	return Brabus{
		ctx:    ctx,
		stop:   cancelFunc,
		Nats:   nats,
		Config: config,
	}
}
