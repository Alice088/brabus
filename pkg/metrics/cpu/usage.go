package cpu

import (
	"fmt"
	"github.com/c9s/goprocinfo/linux"
	"strconv"
	"time"
)

func Usage() ([]string, error) {
	var usages []string

	initialStat, err := linux.ReadStat("/proc/stat")
	if err != nil {
		return usages, fmt.Errorf("cannot read stat(1): %v", err)
	}

	time.Sleep(3 * time.Second)

	newStat, err := linux.ReadStat("/proc/stat")
	if err != nil {
		return usages, fmt.Errorf("cannot read stat(2): %v", err)
	}

	for i, initialCPU := range initialStat.CPUStats {
		newCPU := newStat.CPUStats[i]

		idleTimeDiff := newCPU.Idle - initialCPU.Idle
		totalTimeDiff := TotalWorking(newCPU) - TotalWorking(initialCPU)

		if totalTimeDiff == 0 {
			continue
		}

		cpuUsage := 100.0 * float64(totalTimeDiff-idleTimeDiff) / float64(totalTimeDiff)
		usages = append(usages, strconv.FormatFloat(cpuUsage, 'f', 2, 32))
	}

	return usages, nil
}
