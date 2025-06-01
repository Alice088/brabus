package brabus

import (
	"github.com/rs/zerolog"
	"time"
)

func (brabus *Brabus) Scan(logger *zerolog.Logger) {
	ticker := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-ticker.C:
			brabus.ProcessMetrics(logger)
		case <-brabus.ctx.Done():
			return
		}
	}
}
