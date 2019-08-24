package cacher_test

import (
	"fmt"

	"github.com/y-yagi/cacher"
)

func ExampleFileStore_Read() {
	c := cacher.WithFileStore("/tmp/")
	c.Write("cache-key", []byte("value"), cacher.Forever)

	value, _ := c.Read("cache-key")
	fmt.Print(string(value))
	// Output: value
}

func ExampleFileStore_Delete() {
	c := cacher.WithFileStore("/tmp/")
	c.Write("cache-key", []byte("value"), cacher.Forever)

	c.Delete("cache-key")
	value, _ := c.Read("cache-key")
	fmt.Print(string(value))
	// Output:
}
