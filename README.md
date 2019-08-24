# Cacher

[![GoDoc](https://godoc.org/github.com/y-yagi/cacher?status.svg)](https://godoc.org/github.com/y-yagi/cacher)
[![Go Report Card](https://goreportcard.com/badge/github.com/y-yagi/cacher)](https://goreportcard.com/report/github.com/y-yagi/cacher)
[![Build Status](https://circleci.com/gh/y-yagi/cacher.svg?style=svg)](https://circleci.com/gh/y-yagi/cacher)

Cacher is a simple cache library. Provides cache system using file store by default.

Example:

```go
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
```

## Other stores

* Redis Store: [y-yagi/cacherredis](https://github.com/y-yagi/cacherredis)
