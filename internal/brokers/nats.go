package brokers

import (
	"log"

	"github.com/nats-io/nats.go"
)

type NATSBroker interface {
	Connect()
	Publish(subject string, message []byte) error
	Subscribe(subject string, handler func(msg []byte)) error
	Close()
}

type NATS struct {
	Url        string
	Connection *nats.Conn
}

func New(url ...string) *NATS {
	if len(url) == 0 {
		return &NATS{Url: nats.DefaultURL}
	}
	return &NATS{Url: url[0]}
}

func (b *NATS) Connect() {
	nc, err := nats.Connect(b.Url)
	if err != nil {
		log.Fatalf("Error during connect to NATS: %v", err)
	}
	b.Connection = nc
	log.Println("Successful connect to NATS")
}

func (b *NATS) Publish(subject string, message []byte) error {
	return b.Connection.Publish(subject, message)
}

func (b *NATS) Subscribe(subject string, handler func(m nats.Msg)) error {
	_, err := b.Connection.Subscribe(subject, func(m *nats.Msg) {
	})
	return err
}

func (b *NATS) Close() {
	b.Connection.Close()
	log.Println("Connection with NATS was closed")
}
