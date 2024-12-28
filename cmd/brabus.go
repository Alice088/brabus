package main

import (
	_ "brabus/internal/brokers"
	"brabus/internal/metrics"
	"brabus/internal/metrics/cpu"
	"brabus/internal/metrics/disk"
	"brabus/internal/metrics/ram"
	"brabus/internal/yaml"
	"fmt"
	_ "fmt"
	"github.com/mailru/easyjson"
	natsPkg "github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	nc, err := natsPkg.Connect(natsPkg.DefaultURL)
	defer nc.Close()
	if err != nil {
		log.Fatal(err)
	}

	globalConf := yaml.UnmarshalGlobalConfig()
	failLimit := globalConf.Limits.FailLimit
	failCount := 0

	for {
		time.Sleep(2 * time.Second)
		CPUUsage := cpu.GetCPUUsage()
		RAMUsage := ram.GetRAMUsage()
		DiskSpace := disk.GetDiskFreeSpace()
		DiskUsage := disk.GetDiskUsage()

		someStruct := &metrics.Metrics{
			CPUUsage:  CPUUsage,
			RAMUsage:  RAMUsage,
			DiskSpace: DiskSpace,
			DiskUsage: DiskUsage,
		}
		rawBytes, err := easyjson.Marshal(someStruct)

		if err != nil {
			failCount++
			if failCount > failLimit {
				log.Fatalf("Fatal error marshalling metrics: %v", err)
			}

			log.Printf("Error during marshalling metrics: %v\n", err)
			continue
		}

		_, err = nc.Subscribe("metrics", func(msg *natsPkg.Msg) {
			fmt.Printf("Received a message: %s\n", string(msg.Data))
		})
		if err != nil {
			return
		}

		err = nc.Publish("metrics", rawBytes)
		if err != nil {
			failCount++
			if failCount > failLimit {
				log.Fatalf("Fatal error during publish metrics %v", err)
			}

			log.Printf("Error during publish metrics: %v\n", err)
		}
	}
}
