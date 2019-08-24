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

	value := []byte("dummy")
	cacher := WithFileStore(tempDir)
	got, _ := cacher.Read("cacher-test")
	if got != nil {
		t.Fatalf("want nil, got %q", got)
	}

	cacher.Write("cacher-test", value, Forever)
	got, _ = cacher.Read("cacher-test")
	if string(got) != string(value) {
		t.Fatalf("want %q, got %q", value, got)
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

	value := []byte("dummy")
	cacher := WithFileStore(tempDir)

	cacher.Write("cacher-test", value, 1*time.Second)

	time.Sleep(2 * time.Second)

	got, _ := cacher.Read("cacher-test")
	if got != nil {
		t.Fatalf("want nil, got %q", got)
	}
}

func TestFileStoreCleanup(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "cacher-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	c := WithFileStore(tempDir)
	c.Write("foo", []byte("dummy"), Forever)
	c.Write("bar", []byte("dummy"), Forever)
	c.Write("baz", []byte("dummy"), Forever)
	c.Cleanup()

	keys := []string{"foo", "bar", "baz"}
	for _, key := range keys {
		if got, _ := c.Read(key); got != nil {
			t.Fatalf("want nil, got %q", got)
		}
	}
}
