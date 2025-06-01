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
func Collect() (*dto.Metrics, error) {
	ramUsage, err := ram.Usage()
	if err != nil {
		return nil, err
	}

	cpuUsage, err := cpu.Usage()
	if err != nil {
		return nil, err
	}

	diskSpace, err := disk.Space()
	if err != nil {
		return nil, err
	}

	diskUsage, err := disk.Usage()
	if err != nil {
		return nil, err
	}

	return &dto.Metrics{
		CPU: dto.CPU{
			Usage: cpuUsage,
		},
		RAM: dto.RAM{
			Usage: ramUsage,
		},
		Disk: dto.Disk{
			Space: diskSpace,
			Usage: diskUsage,
		},
	}, nil
}
