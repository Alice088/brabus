package badger

import (
	"brabus/internal/app/banana/storage/badger/iterarators"
	"brabus/pkg/dto"
	"github.com/dgraph-io/badger/v4"
	"github.com/mailru/easyjson"
	"github.com/pkg/errors"
	"strconv"
)

type Percent = float64

var prefix = []byte("cpu:")
var seek = []byte("cpu:~")

func GetCPU(db *badger.DB) ([][]Percent, error) {
	cpus := make([][]Percent, 0)

	err := db.View(func(txn *badger.Txn) error {
		it := iterarators.Reverse(txn, 3)
		defer it.Close()

		count := 0

		for it.Seek(seek); it.ValidForPrefix(prefix) && count < 4; it.Next() {
			item := it.Item()
			itemSlice := make([]Percent, 0)

			var CPU dto.CPU
			err := easyjson.Unmarshal([]byte(item.String()), &CPU)
			if err != nil {
				return errors.Wrap(err, "fail to unmarshal CPU")
			}

			for _, v := range CPU.Usage {
				f, _ := strconv.ParseFloat(v, 64)
				itemSlice = append(itemSlice, f)
				return nil
			}

			cpus = append(cpus, itemSlice)
		}

		return nil
	})

	if err != nil {
		return cpus, errors.Wrap(err, "fail to get CPU")
	}

	if len(cpus) == 0 {
		return cpus, errors.New("cpu not found")
	}

	return cpus, nil
}
