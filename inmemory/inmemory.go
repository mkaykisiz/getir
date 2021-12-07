package inmemory

import (
	"errors"
	"sync"
)

var (
	// KeyNotFoundErr for not founded key error.
	KeyNotFoundErr = errors.New("key not found")
)

// DB : data struct
type DB struct {
	*sync.Map
}

// New returns db.
func New() *DB {
	return &DB{
		&sync.Map{},
	}
}

// Write is set new key.
func (db *DB) Write(key, value string) {
	db.Store(key, value)
}

// Read is get value by key.
func (db *DB) Read(key string) (val string, err error) {
	value, ok := db.Load(key)
	if !ok {
		return "", KeyNotFoundErr
	}

	return value.(string), nil
}
