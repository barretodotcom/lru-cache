package cache

import (
	"time"
)

type Item struct {
	key   string
	value any
	time  time.Time
}

type Cache struct {
	maxCount int
	items    map[string]*Item
}

func New(maxCount int) *Cache {
	return &Cache{
		maxCount: maxCount,
		items:    make(map[string]*Item),
	}
}

func (c *Cache) Get(key string) *Item {
	item, ok := c.items[key]
	if !ok {
		return nil
	}

	c.items[key].time = time.Now()

	return item
}

func (c *Cache) Set(key string, value any) {
	if len(c.items) == c.maxCount {
		oldestTime := time.Now()
		oldestKey := ""
		for i, item := range c.items {
			if item.time.Before(oldestTime) {
				oldestTime = item.time
				oldestKey = i
			}
		}
		delete(c.items, oldestKey)
	}

	c.items[key] = &Item{key: key, value: value, time: time.Now()}
	return
}
