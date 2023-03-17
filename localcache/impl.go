package localcache

import (
	"sync"
	"time"
)

const expireTime = 30 * time.Second

type clock interface {
	Now() time.Time
}

type realClock struct{}

func (realClock) Now() time.Time {
	return time.Now()
}

type cache struct {
	data   map[string]interface{}
	expiry map[string]time.Time
	mutex  sync.RWMutex
	clock  clock
}

func New(c clock) Cache {
	if c == nil {
		c = realClock{}
	}
	return &cache{
		data:   make(map[string]interface{}),
		expiry: make(map[string]time.Time),
		clock:  c,
	}
}

// Get returns the value for the given key, or nil if the key is not found or expired
func (c *cache) Get(key string) interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	value, ok := c.data[key]
	expireTime := c.expiry[key]

	if !ok {
		return nil
	}
	if expireTime.Before(c.clock.Now()) {
		delete(c.data, key)
		delete(c.expiry, key)
		return nil
	}
	return value
}

// Set sets the value for the given key
func (c *cache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = value
	c.expiry[key] = c.clock.Now().Add(expireTime)
}
