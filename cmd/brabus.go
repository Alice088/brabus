package main

import (
	_ "brabus/internal/brokers"
	"fmt"
	nats2 "github.com/nats-io/nats.go"
	"log"
)

func main() {
	nc, err := nats2.Connect(nats2.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	err = nc.Publish("foo", []byte("Hello World"))
	if err != nil {
		log.Fatal(err)
	}

	// Simple Async Subscriber
	_, err = nc.Subscribe("foo", func(m *nats2.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	//globalConf := yaml.UnmarshalGlobalConfig()
	//failLimit := globalConf.Limits.FailLimit
	//failCount := 0
	//
	//for {
	//	time.Sleep(2 * time.Second)
	//	CPUUsage := cpu.GetCPUUsage()
	//	RAMUsage := ram.GetRAMUsage()
	//	DiskSpace := disk.GetDiskFreeSpace()
	//	DiskUsage := disk.GetDiskUsage()
	//
	//	someStruct := &metrics.Metrics{
	//		CPUUsage:  CPUUsage,
	//		RAMUsage:  RAMUsage,
	//		DiskSpace: DiskSpace,
	//		DiskUsage: DiskUsage,
	//	}
	//	rawBytes, err := easyjson.Marshal(someStruct)
	//
	//	if err != nil {
	//		failCount++
	//		if failCount > failLimit {
	//			log.Fatalf("Fatal error marshalling metrics: %v", err)
	//		}
	//
	//		log.Printf("Error during marshalling metrics: %v\n", err)
	//		continue
	//	}
	//
	//	err = nats.Publish("metrics", rawBytes)
	//	if err != nil {
	//		failCount++
	//		if failCount > failLimit {
	//			log.Fatalf("Fatal error during publish metrics %v", err)
	//		}
	//
	//		log.Printf("Error during publish metrics: %v\n", err)
	//	}
	//}
}
