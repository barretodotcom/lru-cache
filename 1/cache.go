package cache

import "time"

type Item struct {
	key   string
	value any
	time  time.Time
}

type LRUCache struct {
	maxCount int
	items    []*Item
}

func New(maxCount int) *LRUCache {
	return &LRUCache{
		maxCount: maxCount,
	}
}

func (l *LRUCache) Get(key string) any {
	for i := range l.items {
		if l.items[i].key == key {
			l.items[i].time = time.Now()
			return l.items[i].value
		}
	}

	return -1
}

func (l *LRUCache) Set(key string, value any) {
	if l.maxCount == len(l.items) {
		oldestKey := -1
		oldestTime := time.Now()
		for i, item := range l.items {
			if item.time.Before(oldestTime) {
				oldestKey = i
				oldestTime = item.time
			}
		}

		l.items[oldestKey].key = key
		l.items[oldestKey].value = value
		return
	}

	l.items = append(l.items, &Item{key: key, value: value})
}
