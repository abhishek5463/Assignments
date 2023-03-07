package cache

import (
	"assignment-2/cache"
	"container/list"
)

type Pair struct {
	key   int
	value interface{}
}

type FifoCache struct {
	cap int
	mp  map[int]*list.Element
	l   *list.List
	
}

func CreateNewFifoCache(size int) cache.Cache {
	return &FifoCache{
		cap: size,
		mp:  make(map[int]*list.Element),
		l:   new(list.List),
	}
}

func (c *FifoCache) Get(key int) interface{} {
	curr, pre := c.mp[key]
	if !pre {
		return "-1"
	}

	v := curr.Value.(Pair).value
	//c.l.MoveToFront(curr)
	return v
}

func (c *FifoCache) Insert(key int, value interface{}) {
	curr, pre := c.mp[key]
	if !pre {
		if len(c.mp) >= c.cap {
			delete(c.mp, Pair(c.l.Back().Value.(Pair)).key)
			c.l.Remove(c.l.Back())
		}
		c.l.PushFront(Pair{key, value})
		c.mp[key] = c.l.Front()
	} else {
		curr.Value = value
	}

}
