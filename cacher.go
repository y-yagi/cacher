package cacher

import (
	"sync"
	"time"
)

// Cacher is a type for cache.
type Cacher struct {
	Store Store
	mu    sync.Mutex
}

// Store is a interface for store.
type Store interface {
	Read(key string) ([]byte, error)
	Write(key string, value []byte, d time.Duration) error
	Delete(key string) error
}

const (
	// Forever use to cache never expire.
	Forever time.Duration = -1
)

// WithFileStore create a new Cache with FileStore.
func WithFileStore(path string) *Cacher {
	cache := &Cacher{}
	cache.Store = &FileStore{path: path}
	return cache
}

// Read cache.
func (c *Cacher) Read(key string) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Store.Read(key)
}

// Write create a new cache.
func (c *Cacher) Write(key string, value []byte, d time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Store.Write(key, value, d)
}

// Delete delete cache.
func (c *Cacher) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Store.Delete(key)
}
