package cpu

import (
	"brabus/pkg/dto"
	"fmt"
)

func CheckAnomalies(cpu dto.CPU, anomalies *[]string) {
	totalUsed := cpu.Average.User + cpu.Average.System
	if totalUsed > 90.0 {
		*anomalies = append(*anomalies, fmt.Sprintf("High CPU usage: %.2f%%", float64(totalUsed)))
	}

	if cpu.Average.IOWait > 10.0 {
		*anomalies = append(*anomalies, fmt.Sprintf("High IOWait: %.2f%%", cpu.Average.IOWait))
	}
}
