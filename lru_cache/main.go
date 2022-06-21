package main

import (
	"container/list"
)

type KeyPair struct {
	key int
	value int
}

type LRUCache struct {
	capacity int
	list *list.List
	elements map[int]*list.Element
}

// key not in the map === -1
// return the value and move it to the front of the list
func (cache *LRUCache) Get(key int) int {
	if node, ok := cache.elements[key]; ok {
		value := node.Value.(*list.Element).Value.(KeyPair).value
		cache.list.MoveToFront(node)
		return value
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if node, ok := cache.elements[key]; ok {
		cache.list.MoveToFront(node)
		node.Value.(*list.Element).Value = KeyPair{key: key, value: value}
	} else {
		if cache.list.Len() == cache.capacity {
			index := cache.list.Back().Value(*list.Element).Value.(KeyPair).key
			delete(cache.elements, index)
			cache.list.Remove(cache.list.Back())
		}
	}

	node := &list.Element{
		Value: KeyPair{
			key: key,
			value: value,
		},
	}

	pointer := cache.list.PushFront(node)
	cache.elements[key] = pointer
}

func (cache *LRUCache) RecentlyUsed() interface{} {
	return cache.list.Front().Value.(*list.Element).Value(KeyPair).value
}

func (cache *LRUCache) Remove(key int) {
	if node, ok := cache.elements[key]; ok {
		delete(cache.elements, key)
		cache.list.Remove(node)
	}
}

func main() {
	
}