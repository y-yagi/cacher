package cacher

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestFileStore(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "cacher-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	data := []byte("dummy")
	cacher := WithFileStore(tempDir)
	got, _ := cacher.Read("cacher-test")
	if got != nil {
		t.Fatalf("want nil, got %q", got)
	}

	cacher.Write("cacher-test", data, Forever)
	got, _ = cacher.Read("cacher-test")
	if string(got) != string(data) {
		t.Fatalf("want %q, got %q", data, got)
	}

	cacher.Delete("cacher-test")
	got, _ = cacher.Read("cacher-test")
	if got != nil {
		t.Fatalf("want nil, got %q", got)
	}
}

func TestFileStoreWithExpired(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "cacher-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	data := []byte("dummy")
	cacher := WithFileStore(tempDir)

	cacher.Write("cacher-test", data, 1*time.Second)

	time.Sleep(2 * time.Second)

	got, _ := cacher.Read("cacher-test")
	if got != nil {
		t.Fatalf("want nil, got %q", got)
	}
}
