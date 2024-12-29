package metrics

import (
	"brabus/internal/DTO"
	"brabus/pkg/metrics/cpu"
	"brabus/pkg/metrics/disk"
	"brabus/pkg/metrics/ram"
)

func CollectMetrics() *DTO.Metrics {
	metricz := new(DTO.Metrics)

	metricz.CPUUsage = cpu.GetCPUUsage()
	metricz.RAMUsage = ram.GetRAMUsage()
	metricz.DiskSpace = disk.GetDiskFreeSpace()
	metricz.DiskUsage = disk.GetDiskUsage()
	return metricz
}
