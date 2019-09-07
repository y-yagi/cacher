package cacher

import (
	"testing"
)

func TestNilStore(t *testing.T) {
	cacher := WithNilStore()

	value := []byte("dummy")
	got, _ := cacher.Read("cacher-test")
	if got != nil {
		t.Fatalf("want nil, got %q", got)
	}

	cacher.Write("cacher-test", value, Forever)
	got, _ = cacher.Read("cacher-test")
	if got != nil {
		t.Fatalf("want %q, got %q", value, got)
	}

	cacher.Delete("cacher-test")
	got, _ = cacher.Read("cacher-test")
	if got != nil {
		t.Fatalf("want nil, got %q", got)
	}
}
