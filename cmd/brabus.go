package main

import (
	"brabus/internal/metrics/cpu"
	disk2 "brabus/internal/metrics/disk"
	"brabus/internal/metrics/ram"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Error during connect to NATS: %v", err)
	}
	defer nc.Close()
	log.Println("NATS was connected to brabus")

	for {
		time.Sleep(2 * time.Second)
		cpU, _ := cpu.GetCPUUsage()
		rmU, _ := ram.GetRAMUsage()
		dcS := disk2.GetDiskFreeSpace()
		dcU := disk2.GetDiskUsage()

		fmt.Println(cpU, rmU, dcS, dcU)
	}
}
