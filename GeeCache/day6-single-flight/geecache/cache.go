package geecache

import (
	"fmt"
	"github.com/qinchenfeng/HelloGoSevenDays/GeeCache/day6-single-flight/geecache/lru"
	"sync"
)

type cache struct {
	mu         sync.Mutex
	lru        *lru.Cache
	cacheBytes int64
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

var getCount int

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	fmt.Printf("get count:%d\n", getCount)
	getCount++
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}

	if v, ok := c.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}
