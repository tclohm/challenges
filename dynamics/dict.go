package main

import (
	"fmt"
	"sync"
)

type Key string
type Value string

type Dictionary struct {
	elements map[Key]Value
	lock sync.RWMutex
}

func (dict *Dictionary) Put(key Key, value Value) {
	dict.lock.Lock()
	defer dict.lock.Unlock()

	if dict.elements == nil {
		dict.elements = make(map[Key]Value)
	}

	dict.elements[key] = value
}

func (dict *Dictionary) Remove(key Key) bool {
	dict.lock.Lock()
	defer dict.lock.Unlock()

	_, exists := dict.elements[key]

	if exists {
		delete(dict.elements, key)
	}

	return exists
}

func (dict *Dictionary) Contains(key Key) bool {
	dict.lock.RLock()
	defer dict.lock.RUnlock()

	_, exists := dict.elements[key]
	return exists
}

func (dict *Dictionary) Find(key Key) Value {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	return dict.elements[key]
}

func (dict *Dictionary) Reset() {
	dict.lock.Lock()
	defer dict.lock.Unlock()
	dict.elements = make(map[Key]Value)
}

func (dict *Dictionary) Length() int {
	dict.lock.RLock()
	defer dict.lock.RUnlock()
	return len(dict.elements)
}

func (dict *Dictionary) GetKeys() []Key {
	dict.lock.RLock()
	defer dict.lock.RUnlock()

	var keys = []Key{}

	for key := range dict.elements {
		keys = append(keys, key)
	}

	return keys
}

func (dict *Dictionary) GetValues() []Value {
	dict.lock.RLock()
	defer dict.lock.RUnlock()

	var values = []Value{}

	for key := range dict.elements {
		values = append(values,  dict.elements[key])
	}

	return values
}

func main() {
	var d *Dictionary = &Dictionary{}
	d.Put("Hello", "world")
	d.Put("Soccer", "Ball")
	fmt.Println(d)
}