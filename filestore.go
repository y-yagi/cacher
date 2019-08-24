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

// FileStore is a type for file store.
type FileStore struct {
	path string
}

type entry struct {
	Value      []byte
	Expiration int64
}

const fileprefix = "cacher-"

// Read reads cache from a file store.
func (fs *FileStore) Read(key string) ([]byte, error) {
	filename := fs.filename(key)
	if !osext.IsExist(filename) {
		return nil, nil
	}

	e, err := fs.readFile(filename)
	if err != nil {
		return nil, err
	}

	if e.expired() {
		return nil, nil
	}

	return e.Value, nil
}

// Write stores data to a file store.
func (fs *FileStore) Write(key string, value []byte, d time.Duration) error {
	e := &entry{Value: value}
	if d > 0 {
		e.Expiration = time.Now().Add(d).UnixNano()
	}

	return ioutil.WriteFile(fs.filename(key), fs.encode(e), 0644)
}

// Delete deletes data from a file store.
func (fs *FileStore) Delete(key string) error {
	filename := fs.filename(key)
	if !osext.IsExist(filename) {
		return nil
	}

	return os.Remove(filename)
}

// Cleanup deletes the expired cache.
func (fs *FileStore) Cleanup() error {
	matches, err := filepath.Glob(filepath.Join(fs.path, fileprefix+"*"))
	if err != nil {
		return err
	}

	for _, match := range matches {
		e, err := fs.readFile(match)
		if err != nil {
			return err
		}

		if e.expired() {
			if err = os.Remove(match); err != nil {
				return err
			}
		}
	}

	return nil
}

func (fs *FileStore) readFile(filename string) (*entry, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return fs.decode(b), err
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
	return filepath.Join(fs.path, fileprefix+key)
}

func (e *entry) expired() bool {
	if e.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > e.Expiration
}
