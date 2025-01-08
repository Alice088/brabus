package metrics

import (
	"brabus/internal/DTO"
	"brabus/pkg/metrics/cpu"
	"brabus/pkg/metrics/disk"
	"brabus/pkg/metrics/ram"
)

func CollectMetrics() *DTO.Metrics {
	metrics := new(DTO.Metrics)

	metrics.CPU.Usage = cpu.Usage()
	metrics.RAM.Usage = ram.Usage()
	metrics.Disk.Space = disk.FreeSpace()
	metrics.Disk.Usage = disk.Usage()
	return metrics
}
