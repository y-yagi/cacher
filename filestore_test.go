package cacher

import (
	"io/ioutil"
	"os"
	"testing"
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

	cacher.Write("cacher-test", data)
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
