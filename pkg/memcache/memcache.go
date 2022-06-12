package memcache

import (
	"fmt"
	"sync"
	"time"
)

var ErrorKeyNotFound = fmt.Errorf("key not found")

type memCache struct {
	sync.Mutex
	items map[string]item
}

type item struct {
	value          interface{}
	expirationDate int64
}

func NewMemCache() *memCache {
	return &memCache{
		items: make(map[string]item),
	}
}

func (m *memCache) Set(key string, value interface{}, ttl time.Duration) {
	m.Lock()
	defer m.Unlock()

	m.items[key] = item{
		value:          value,
		expirationDate: time.Now().Add(ttl).Unix(),
	}
}

func (m *memCache) Get(key string) (interface{}, bool) {
	m.Lock()
	defer m.Unlock()

	item, ok := m.items[key]

	if !ok {
		return nil, ok
	}

	if time.Now().Unix() > item.expirationDate {
		delete(m.items, key)

		return nil, false
	}

	return item.value, true
}

func (m *memCache) Delete(key string) error {
	m.Lock()
	defer m.Unlock()

	_, ok := m.items[key]

	if !ok {
		return ErrorKeyNotFound
	}

	delete(m.items, key)

	return nil
}
