package cpu

import (
	"brabus/pkg/dto"
	"fmt"
)

func CheckLoad(cpu dto.CPU, anomalies *[]string) {
	cores := float64(cpu.Cores)

	thresholds := []struct {
		value    float64
		period   string
		severity string
	}{
		{cpu.Load1, "1min", "warning"},
		{cpu.Load5, "5min", "warning"},
		{cpu.Load15, "15min", "critical"},
	}

	for _, t := range thresholds {
		if t.value > cores*0.9 { // 90% load
			*anomalies = append(*anomalies,
				fmt.Sprintf("[%s] CRITICAL load (%s): %.2f%% (cores: %d)",
					t.severity, t.period, t.value, cpu.Cores))
		} else if t.value > cores*0.7 { // 70% load
			*anomalies = append(*anomalies,
				fmt.Sprintf("[%s] WARNING load (%s): %.2f%% (cores: %d)",
					t.severity, t.period, t.value, cpu.Cores))
		}
	}

	if cpu.Load1 > cpu.Load15*1.5 {
		*anomalies = append(*anomalies, fmt.Sprintf("[critical] Spike detected!: %.2f%%(1min) > %.2f%%(15min)", cpu.Load1, cpu.Load15*1.5))
	}
}
