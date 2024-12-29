package main

import (
	_ "brabus/internal/brokers"
	"brabus/internal/service/brabus"
)

func main() {
	brabus.Run()
}
