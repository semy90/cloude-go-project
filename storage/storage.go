package storage

import "errors"

var storage map[string]string

var (
	ErrorNoSuchKey = errors.New("no such a key")
)

func Put(key string, value string) error {
	storage[key] = value
	return nil
}

func Get(key string) (string, error) {
	value, ok := storage[key]
	if !ok {
		return "", ErrorNoSuchKey
	}
	return value, nil
}

func Delete(key string) {
	delete(storage, key)
}
