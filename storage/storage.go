package storage

import (
	"errors"
	"sync"
)

var storage = struct {
	sync.RWMutex
	m map[string]string
}{
	m: make(map[string]string),
}

var (
	ErrorNoSuchKey = errors.New("no such a key")
)

func Put(key string, value string) error {
	const op = "storage.Put"
	storage.Lock()
	storage.m[key] = value
	storage.Unlock()
	return nil
}

func Get(key string) (string, error) {
	const op = "storage.Get"

	storage.RLock()
	value, ok := storage.m[key]
	storage.RUnlock()
	if !ok {
		return "", ErrorNoSuchKey
	}
	return value, nil
}

func Delete(key string) {
	const op = "storage.Delete"

	storage.Lock()
	delete(storage.m, key)
	storage.Unlock()
}
