package main

import (
	"fmt"

	"training/localcache"
)

func main() {
	cache := localcache.New(nil)

	// set a key/value pair in the cache
	cache.Set("foo", 123)

	// retrieve a value from the cache by key
	value := cache.Get("foo")

	fmt.Println(value) // output: "bar"
}
