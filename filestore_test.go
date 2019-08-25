package cacher

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

var tempDir string

func TestMain(m *testing.M) {
	var err error
	tempDir, err = ioutil.TempDir("", "cacher-test")
	if err != nil {
		panic(err)
	}

	code := m.Run()

	os.RemoveAll(tempDir)
	os.Exit(code)
}

func TestFileStore(t *testing.T) {
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

func TestFileStoreWith_Expired(t *testing.T) {
	value := []byte("dummy")
	cacher := WithFileStore(tempDir)

	cacher.Write("cacher-test", value, 1*time.Second)

	time.Sleep(2 * time.Second)

	got, _ := cacher.Read("cacher-test")
	if got != nil {
		t.Fatalf("want nil, got %q", got)
	}
}

func TestFileStore_Cleanup(t *testing.T) {
	c := WithFileStore(tempDir)
	c.Write("foo", []byte("foo"), Forever)
	c.Write("bar", []byte("bar"), 1*time.Second)
	c.Write("baz", []byte("baz"), Forever)

	time.Sleep(2 * time.Second)
	c.Cleanup()

	keys := map[string]string{"foo": "foo", "bar": "", "baz": "baz"}
	for key, want := range keys {
		if got, _ := c.Read(key); string(got) != want {
			t.Fatalf("want %q, got %q", want, got)
		}
	}
}

func TestFileStore_Exist(t *testing.T) {
	c := WithFileStore(tempDir)
	if c.Exist("not-exist") {
		t.Fatalf("want false, got true")
	}

	c.Write("exist", []byte("foo"), Forever)
	if !c.Exist("exist") {
		t.Fatalf("want true, got false")
	}
}
