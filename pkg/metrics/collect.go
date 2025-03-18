package metrics

import (
	"brabus/pkg/dto"
	"brabus/pkg/metrics/cpu"
	"brabus/pkg/metrics/disk"
	"brabus/pkg/metrics/ram"
)

func Collect() dto.Metrics {
	return dto.Metrics{
		CPU: dto.CPU{
			Usage: cpu.Usage(),
		},
		RAM: dto.RAM{
			Usage: ram.Usage(),
		},
		Disk: dto.Disk{
			Space: disk.Space(),
			Usage: disk.Usage(),
		},
	}
}
