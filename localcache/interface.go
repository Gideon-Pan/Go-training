package localcache

type Cache interface {
	Get(key string) interface{}
	Set(key string, value interface{})
}
