package metrics

import (
	"brabus/pkg/dto"
	"brabus/pkg/metrics/cpu"
	"brabus/pkg/metrics/disk"
	"brabus/pkg/metrics/ram"
)

func Collect() dto.Metrics {
	metrics := dto.Metrics{}

	metrics.CPU.Usage = cpu.Usage()
	metrics.RAM.Usage = ram.Usage()
	metrics.Disk.Space = disk.FreeSpace()
	metrics.Disk.Usage = disk.Usage()
	return metrics
}
