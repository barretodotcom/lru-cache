package cache

type Cache struct {
	dict     map[string]*Item
	maxCount int
	items    *Item
	head     *Item
}

type Item struct {
	key   string
	value any
	prev  *Item
	next  *Item
}

func New(maxCount int) *Cache {
	return &Cache{
		maxCount: maxCount,
		dict:     make(map[string]*Item),
	}
}

func (c *Cache) Get(key string) any {
	item, ok := c.dict[key]
	if !ok {
		return -1
	}

	if item == c.head {
		return item.value
	}

	if item.prev != nil {
		item.prev.next = item.next
	}

	if item == c.items {
		c.items = item.next
	}

	c.head.next = item
	item.prev = c.head
	c.head = item
	c.head.next = nil

	return item.value
}

func (c *Cache) Set(key string, value any) {
	if len(c.dict) == c.maxCount {
		item := c.items

		delete(c.dict, item.key)
		item.key = key
		item.value = value
		c.dict[key] = item
		return
	}
	newItem := &Item{key: key, value: value}
	newItem.next = c.items
	c.items = newItem
	c.dict[key] = newItem

	if c.items.next == nil {
		c.head = c.items
	}
}
