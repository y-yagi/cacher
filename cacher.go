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
	Cleanup() error
	Exist(key string) bool
}

const (
	// Forever means cache never expire.
	Forever time.Duration = -1
)

// WithFileStore create a new Cacher with a file store.
func WithFileStore(path string) *Cacher {
	cache := &Cacher{}
	cache.Store = &FileStore{path: path}
	return cache
}

// Read reads cache from a store.
func (c *Cacher) Read(key string) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Store.Read(key)
}

// Write stores data to a store.
func (c *Cacher) Write(key string, value []byte, d time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Store.Write(key, value, d)
}

// Delete deletes data from a file store.
func (c *Cacher) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Store.Delete(key)
}

// Cleanup deletes expired cache.
func (c *Cacher) Cleanup() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Store.Cleanup()
}

// Exist check the cache exists or not.
func (c *Cacher) Exist(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Store.Exist(key)
}
