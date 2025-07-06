package analyze

import (
	"brabus/pkg/dto"
	"github.com/rs/zerolog"
)

func Metrics(metrics dto.Metrics, logger zerolog.Logger) {
	CPU(metrics, logger)
	Disk(metrics)
	RAM(metrics)
}
