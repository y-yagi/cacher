package cacher

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/y-yagi/goext/osext"
)

// FileStore is a type for FileStore.
type FileStore struct {
	path string
}

// NewFileStore create a new FileStore.
func NewFileStore(path string) *FileStore {
	fs := &FileStore{path: path}
	return fs
}

// Read cache.
func (fs *FileStore) Read(key string) ([]byte, error) {
	file := filepath.Join(fs.path, key)
	if !osext.IsExist(file) {
		return nil, nil
	}
	return ioutil.ReadFile(file)
}

// Write create a new cache.
func (fs *FileStore) Write(key string, data []byte) error {
	file := filepath.Join(fs.path, key)
	return ioutil.WriteFile(file, data, 0644)
}

// Delete delete cache.
func (fs *FileStore) Delete(key string) error {
	file := filepath.Join(fs.path, key)
	if !osext.IsExist(file) {
		return nil
	}

	return os.Remove(file)
}
