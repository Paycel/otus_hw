package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type cacheItem struct {
	key   Key
	value any
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem

	mu *sync.Mutex
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mu:       &sync.Mutex{},
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if v, ok := l.items[key]; ok {
		l.items[key].Value.(*cacheItem).value = value
		l.queue.MoveToFront(v)
		return true
	}
	elem := l.queue.PushFront(value)
	elem.Value = &cacheItem{key: key, value: value}
	l.items[key] = elem
	if l.queue.Len() > l.capacity {
		back := l.queue.Back()
		l.queue.Remove(back)
		delete(l.items, back.Value.(*cacheItem).key)
	}
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	v, ok := l.items[key]
	if !ok {
		return nil, false
	}
	l.queue.MoveToFront(v)
	return v.Value.(*cacheItem).value, true
}

func (l *lruCache) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}
