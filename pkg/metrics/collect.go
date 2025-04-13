package metrics

import (
	"brabus/pkg/dto"
	"brabus/pkg/metrics/cpu"
	"brabus/pkg/metrics/disk"
	"brabus/pkg/metrics/ram"
)

// Collect  metrics of a machine
//
//Example:
/*
  "cpu" : {
    "usage" : [ "10.33", "20.33", "9.36", "7.72", "8.67", "13.09", "4.73", "8.67" ]
  },
  "ram" : {
    "usage" : "30.39"
  },
  "disk" : {
    "space" : "182.45",
    "usage" : "44.64"
  }
*/
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
