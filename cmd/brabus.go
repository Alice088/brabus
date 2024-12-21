package main

import (
	"brabus/internal/cpu"
	"fmt"
	"time"
)

func main() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println(cpu.GetCPUUsage())
	}
}
