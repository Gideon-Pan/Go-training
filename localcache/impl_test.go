package localcache

import (
	"testing"
	"time"
)

type mockClock struct {
	now time.Time
}

func (mc mockClock) Now() time.Time {
	return mc.now
}

var mockNow = time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC)

func TestLocalCache(t *testing.T) {
	t.Run("cached word", func(t *testing.T) {
		mc := mockClock{now: mockNow}
		cache := New(mc)
		cache.Set("foo", 123)
		got := cache.Get("foo")
		want := 123
		AssertEqual(t, want, got)
	})

	t.Run("uncached word", func(t *testing.T) {
		mc := mockClock{now: mockNow}
		cache := New(mc)
		cache.Set("foo", 123)
		got := cache.Get("wrong key")
		AssertEqual(t, nil, got)
	})

	t.Run("cached word expired", func(t *testing.T) {
		mc := &mockClock{now: mockNow}
		cache := New(mc)
		cache.Set("foo", 123)
		mc.now = mc.now.Add(expireTime + 1*time.Second)
		got := cache.Get("foo")
		AssertEqual(t, nil, got)
	})

}

func AssertEqual(t *testing.T, want interface{}, got interface{}) {
	t.Helper()
	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
