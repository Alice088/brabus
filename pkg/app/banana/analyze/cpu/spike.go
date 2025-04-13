package cpu

import (
	"brabus/pkg/app/banana/storage/badger"
	"brabus/pkg/dto"
	bad "github.com/dgraph-io/badger/v4"
	"github.com/pkg/errors"
)

func Spike(cpu dto.CPU, db *bad.DB) error {
	oldCpu, err := badger.GetCPU(db)
	if err != nil {
		return errors.Wrap(err, "error analyze spike")
	}

}
