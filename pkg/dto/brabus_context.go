package dto

import "github.com/nats-io/nats.go"

type BrabusContext struct {
	FailCount  *int
	GlobalConf Global
	NC         *nats.Conn
	Quit       chan int
}
