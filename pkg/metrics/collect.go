package metrics

import (
	"brabus/pkg/config"
	"brabus/pkg/dto"
	"brabus/pkg/metrics/cpu"
	"brabus/pkg/metrics/disk"
	"brabus/pkg/metrics/ram"
	"github.com/rs/zerolog"
	"time"
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
func Collect(logger *zerolog.Logger, conf config.Config) (*dto.Metrics, error) {

	t := time.Now()
	ramUsage, err := ram.Usage()
	logger.Debug().Msgf("Duration - [ram.usage]: %s", time.Since(t).String())
	if err != nil {
		return nil, err
	}

	t = time.Now()
	cpuUsage, err := cpu.Usage(conf)
	logger.Debug().Msgf("Duration - [cpu.usage]: %s", time.Since(t).String())
	if err != nil {
		return nil, err
	}

	t = time.Now()
	diskSpace, err := disk.Space()
	logger.Debug().Msgf("Duration - [disk.space]: %s", time.Since(t).String())
	if err != nil {
		return nil, err
	}

	t = time.Now()
	diskUsage, err := disk.Usage()
	logger.Debug().Msgf("Duration - [disk.usage]: %s", time.Since(t).String())
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
