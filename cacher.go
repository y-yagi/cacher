package cacher

import (
	"bytes"
	"encoding/gob"
	"sync"
	"time"
)

// Cacher is a type for cache.
type Cacher struct {
	store Store
	mu    sync.Mutex
}

// Store is a interface for store.
type Store interface {
	Read(key string) ([]byte, error)
	Write(key string, value []byte) error
	Delete(key string) error
}

const (
	// Forever use to cache never expire.
	Forever time.Duration = -1
)

type entry struct {
	Value      []byte
	Expiration int64
}

func (e *entry) expired() bool {
	if e.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > e.Expiration
}

// WithFileStore create a new Cache with FileStore.
func WithFileStore(path string) *Cacher {
	cache := &Cacher{}
	cache.store = &FileStore{path: path}
	return cache
}

// Read cache.
func (c *Cacher) Read(key string) ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	d, err := c.store.Read(key)
	if err != nil {
		return nil, err
	}

	e := decode(d)
	if e.expired() {
		return nil, nil
	}

	return e.Value, nil
}

// Write create a new cache.
func (c *Cacher) Write(key string, value []byte, d time.Duration) error {
	e := &entry{Value: value}
	if d > 0 {
		e.Expiration = time.Now().Add(d).UnixNano()
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	return c.store.Write(key, encode(e))
}

// Delete delete cache.
func (c *Cacher) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.store.Delete(key)
}

func encode(e *entry) []byte {
	buf := bytes.NewBuffer(nil)
	_ = gob.NewEncoder(buf).Encode(e)
	return buf.Bytes()
}

func decode(data []byte) *entry {
	var e entry
	buf := bytes.NewBuffer(data)
	_ = gob.NewDecoder(buf).Decode(&e)
	return &e
}
