package cpu

import (
	"github.com/c9s/goprocinfo/linux"
	"log"
	"strconv"
	"time"
)

func GetCPUUsage() ([]string, error) {
	var CPUUsages []string

	initialStat, err := linux.ReadStat("/proc/stat")
	if err != nil {
		log.Printf("Error during reading /proc/stat: %v\n\n", err)
		return nil, err
	}

	time.Sleep(3 * time.Second)

	newStat, err := linux.ReadStat("/proc/stat")
	if err != nil {
		log.Printf("Error during reading /proc/stat: %v\n\n", err)
		return nil, err
	}

	for i, initialCPU := range initialStat.CPUStats {
		newCPU := newStat.CPUStats[i]

		idleTimeDiff := newCPU.Idle - initialCPU.Idle
		totalTimeDiff := GetTotalCPUWorking(newCPU) - GetTotalCPUWorking(initialCPU)

		if totalTimeDiff == 0 {
			continue
		}

		cpuUsage := 100.0 * float64(totalTimeDiff-idleTimeDiff) / float64(totalTimeDiff)
		CPUUsages = append(CPUUsages, strconv.FormatFloat(cpuUsage, 'f', 2, 32))
	}

	return CPUUsages, nil
}
