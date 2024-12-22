package main

import (
	"brabus/internal/metrics"
	"brabus/internal/metrics/cpu"
	disk2 "brabus/internal/metrics/disk"
	"brabus/internal/metrics/ram"
	"github.com/mailru/easyjson"
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
		CPUUsage, _ := cpu.GetCPUUsage()
		RAMUsage, _ := ram.GetRAMUsage()
		DiskSpace := disk2.GetDiskFreeSpace()
		DiskUsage := disk2.GetDiskUsage()

		someStruct := &metrics.Metrics{
			CPUUsage:  CPUUsage,
			RAMUsage:  RAMUsage,
			DiskSpace: DiskSpace,
			DiskUsage: DiskUsage,
		}
		rawBytes, err := easyjson.Marshal(someStruct)

		if err != nil {
			log.Fatalf("Error during marshalling metrics: %v", err)
		}

		log.Println(string(rawBytes))
	}
}
