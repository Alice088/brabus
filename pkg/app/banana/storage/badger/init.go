package badger

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/pkg/errors"
)

func Init(path ...string) (*badger.DB, error) {
	if len(path) == 0 {
		path = []string{"./tmp/badger"}
	}

	db, err := badger.Open(badger.DefaultOptions(path[0]))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open badger database: %s", path[0])
	}

	return db, nil
}
