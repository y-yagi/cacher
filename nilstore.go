package cacher

import (
	"time"
)

// NilStore is a type for nil store.
type NilStore struct{}

// Read reads cache from a nil store.
func (ns *NilStore) Read(key string) ([]byte, error) {
	return nil, nil
}

// Write stores data to a nil store.
func (ns *NilStore) Write(key string, value []byte, d time.Duration) error {
	return nil
}

// Delete deletes data from a nil store.
func (ns *NilStore) Delete(key string) error {
	return nil
}

// Cleanup deletes the expired cache.
func (ns *NilStore) Cleanup() error {
	return nil
}

// Exist check the cache exists or not.
func (ns *NilStore) Exist(key string) bool {
	return false
}
