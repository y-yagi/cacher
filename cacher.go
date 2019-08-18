package cacher

// Cacher is a type for cache.
type Cacher struct {
	store Store
}

// Store is a interface for store.
type Store interface {
	Read(key string) ([]byte, error)
	Write(key string, value []byte) error
	Delete(key string) error
}

// WithFileStore create a new Cache with FileStore.
func WithFileStore(path string) *Cacher {
	cache := &Cacher{}
	cache.store = &FileStore{path: path}
	return cache
}

// Read cache.
func (c *Cacher) Read(key string) ([]byte, error) {
	return c.store.Read(key)
}

// Write create a new cache.
func (c *Cacher) Write(key string, value []byte) error {
	return c.store.Write(key, value)
}

// Delete delete cache.
func (c *Cacher) Delete(key string) error {
	return c.store.Delete(key)
}
