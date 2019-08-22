package cacher

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/y-yagi/goext/osext"
)

// FileStore is a type for FileStore.
type FileStore struct {
	path string
}

type entry struct {
	Value      []byte
	Expiration int64
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

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	e := fs.decode(b)
	if e.expired() {
		return nil, nil
	}

	return e.Value, nil
}

// Write create a new cache.
func (fs *FileStore) Write(key string, data []byte, d time.Duration) error {
	e := &entry{Value: data}
	if d > 0 {
		e.Expiration = time.Now().Add(d).UnixNano()
	}

	file := filepath.Join(fs.path, key)
	return ioutil.WriteFile(file, fs.encode(e), 0644)
}

// Delete delete cache.
func (fs *FileStore) Delete(key string) error {
	file := filepath.Join(fs.path, key)
	if !osext.IsExist(file) {
		return nil
	}

	return os.Remove(file)
}

func (fs *FileStore) encode(e *entry) []byte {
	buf := bytes.NewBuffer(nil)
	_ = gob.NewEncoder(buf).Encode(e)
	return buf.Bytes()
}

func (fs *FileStore) decode(data []byte) *entry {
	var e entry
	buf := bytes.NewBuffer(data)
	_ = gob.NewDecoder(buf).Decode(&e)
	return &e
}

func (e *entry) expired() bool {
	if e.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > e.Expiration
}
