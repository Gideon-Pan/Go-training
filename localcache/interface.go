package localcache

type LocalCache interface {
	Get(key string) (interface{})
	Set(key string, value interface{})
}