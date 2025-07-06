package brabus

import (
	"context"
	"github.com/rs/zerolog"
	"os"
	"sync"
	"time"
)

type Shutdown struct {
	Ctx context.Context
	Os  chan os.Signal
}

func (brabus *Brabus) Run(logger zerolog.Logger, wg *sync.WaitGroup, shutdown Shutdown) {
	ticker := time.NewTicker(2 * time.Second)

	for {
		select {
		case <-ticker.C:
			wg.Add(1)
			go func() {
				defer wg.Done()
				brabus.CollectMetrics(logger, shutdown.Os)
			}()
			wg.Wait()
		case <-shutdown.Ctx.Done():
			return
		case <-shutdown.Os:
			return
		}
	}
}
