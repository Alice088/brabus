package analyze

import (
	"brabus/pkg/dto"
)

func Metrics(metrics dto.Metrics) {
	AnalyzeCPU(metrics)
	AnalyzeDisk(metrics)
	AnalyzeRAM(metrics)
}
