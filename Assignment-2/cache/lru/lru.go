package cache

import (
	"assignment-2/cache"
	"assignment-2/cache/redis"
	"container/list"
)

type Pair struct {
	key   int
	value interface{}
}

type LRUCache struct {
	cap   int
	mp    map[int]*list.Element
	l     *list.List
	redis cache.Cache
}

// Creating new object of LRU cache
func CreateNewLRUCache(size int) cache.Cache {
	return &LRUCache{
		cap:   size,
		mp:    make(map[int]*list.Element),
		l:     new(list.List),
		redis: &redis.Redis{},
	}
}

func (c *LRUCache) Get(s int) interface{} {
	curr, pre := c.mp[s]
	if !pre {
		return c.redis.Get(s)
	}

	v := curr.Value.(Pair).value
	c.l.MoveToFront(curr)
	return v
}

func (c *LRUCache) Insert(key int, value interface{}) {
	curr, pre := c.mp[key]
	if pre {
		delete(c.mp, Pair(curr.Value.(Pair)).key)
		c.l.Remove(curr)
	}
	if len(c.mp) >= c.cap {
		delete(c.mp, Pair(c.l.Back().Value.(Pair)).key)
		c.l.Remove(c.l.Back())
	}
	c.l.PushFront(Pair{key, value})
	c.mp[key] = c.l.Front()
}
