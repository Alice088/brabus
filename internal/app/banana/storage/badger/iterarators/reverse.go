package iterarators

import "github.com/dgraph-io/badger/v4"

func Reverse(txn *badger.Txn, prefetch ...int) *badger.Iterator {
	opts := badger.DefaultIteratorOptions
	opts.Reverse = true

	if len(prefetch) != 0 {
		opts.PrefetchSize = prefetch[0]
	}

	return txn.NewIterator(opts)
}
