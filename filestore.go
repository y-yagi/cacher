package cacher

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/y-yagi/goext/osext"
)

type FileStore struct {
	path string
}

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
func (fs *FileStore) Write(key string, value []byte) error {
	file := filepath.Join(fs.path, key)
	return ioutil.WriteFile(file, value, 0644)
}

// Delete delete cache.
func (fs *FileStore) Delete(key string) error {
	file := filepath.Join(fs.path, key)
	if !osext.IsExist(file) {
		return nil
	}

	return os.Remove(file)
}
