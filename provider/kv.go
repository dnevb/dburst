package provider

import (
	"encoding/json"

	"github.com/dgraph-io/badger/v3"
)

type KV struct {
	DB *badger.DB
}

func CreateKv(conf *Config) *KV {
	opts := badger.DefaultOptions(conf.Path)
	db, _ := badger.Open(opts.WithLoggingLevel(badger.ERROR))

	return &KV{DB: db}
}

func (kv *KV) View(fn func(txn *badger.Txn) error) error {
	return kv.DB.View(fn)
}

func (kv *KV) Set(key string, value any) error {

	msg, err := json.Marshal(value)
	if err != nil {
		return err
	}

	txn := kv.DB.NewTransaction(true)
	err = txn.Set([]byte(key), msg)
	if err != nil {
		return err
	}

	return txn.Commit()
}

func (kv *KV) Get(key string, dest any) (bool, error) {

	err := kv.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		v, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return json.Unmarshal(v, dest)
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (kv *KV) Delete(key []byte) error {
	return kv.DB.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}

func (kv *KV) Keys() ([][]byte, error) {
	var ks [][]byte
	err := kv.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close() //nolint:errcheck
		for it.Rewind(); it.Valid(); it.Next() {
			ks = append(ks, it.Item().KeyCopy(nil))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ks, nil
}

func (kv *KV) PrefixValues(key string) ([][]byte, error) {
	prefix := []byte(key)
	var values [][]byte

	err := kv.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			v, _ := it.Item().ValueCopy(nil)

			values = append(values, v)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return values, nil
}
