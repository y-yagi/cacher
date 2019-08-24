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

// Read cache.
func (fs *FileStore) Read(key string) ([]byte, error) {
	filename := fs.filename(key)
	if !osext.IsExist(filename) {
		return nil, nil
	}

	b, err := ioutil.ReadFile(filename)
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
func (fs *FileStore) Write(key string, value []byte, d time.Duration) error {
	e := &entry{Value: value}
	if d > 0 {
		e.Expiration = time.Now().Add(d).UnixNano()
	}

	return ioutil.WriteFile(fs.filename(key), fs.encode(e), 0644)
}

// Delete delete cache.
func (fs *FileStore) Delete(key string) error {
	filename := fs.filename(key)
	if !osext.IsExist(filename) {
		return nil
	}

	return os.Remove(filename)
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

func (fs *FileStore) filename(key string) string {
	return filepath.Join(fs.path, "cacher-"+key)
}

func (e *entry) expired() bool {
	if e.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > e.Expiration
}
