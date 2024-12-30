package analyze

import "brabus/internal/DTO"

func ProcessMetrics(metrics *DTO.Metrics) {
	AnalyzeCPU(metrics)
	AnalyzeDisk(metrics)
	AnalyzeRAM(metrics)
}
