/*
The localcache package uses memory to provide a caching mechanism.
*/
package localcache

// Cache will be implemented by localcache/impl
type Cache interface {
	// Get returns the value for the given key, or nil if the key is not found
	Get(key string) interface{}
	// Set sets the value for the given key
	Set(key string, value interface{})
}
