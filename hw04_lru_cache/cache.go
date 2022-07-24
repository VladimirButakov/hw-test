package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
	Len() int
	QueueLen() int
}

type lruCache struct {
	sync.RWMutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	cache.Lock()
	defer cache.Unlock()

	var alreadyExist bool

	if prevValue, ok := cache.items[key]; ok {
		cache.queue.Remove(prevValue)

		alreadyExist = true
	} else {
		alreadyExist = false

		cache.checkCapacity(1)
	}

	cache.items[key] = cache.queue.PushFront(value)

	return alreadyExist
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	cache.Lock()
	defer cache.Unlock()

	if item, ok := cache.items[key]; ok {
		cache.queue.PushFront(item.Value)

		return item.Value, true
	}

	return nil, false
}

func (cache *lruCache) Clear() {
	cache.queue.ClearFrontBack()
	cache.items = make(map[Key]*ListItem)
}

func (cache *lruCache) checkCapacity(extra int) {
	if cache.capacity < (cache.queue.Len() + extra) {
		l := cache.queue

		l.Remove(l.Back())
	}
}

func (cache *lruCache) Len() int {
	return len(cache.items)
}

func (cache *lruCache) QueueLen() int {
	return cache.queue.Len()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
