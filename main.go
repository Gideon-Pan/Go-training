package main

import (
	"fmt"

	"training/localcache"
)

func main() {
	cache := localcache.New(nil)
	cache.Set("foo", 123)
	value := cache.Get("foo")
	fmt.Println(value)
}
