package analyze

import (
	"brabus/internal/app/banana/analyze/cpu"
	"brabus/pkg/dto"
	"fmt"
	"github.com/rs/zerolog"
)

func CPU(metrics dto.Metrics, logger zerolog.Logger) {
	anomalies := new([]string)

	cpu.CheckAnomalies(metrics.CPU, anomalies)
	cpu.CheckLoad(metrics.CPU, anomalies)

	if len(*anomalies) != 0 {
		for i, anomaly := range *anomalies {
			logger.Warn().Msgf("Anomaly[%d]: %s", i+1, anomaly)
		}
		fmt.Println("--------------------------------TICK(2s)--------------------------------------")
	}
}
