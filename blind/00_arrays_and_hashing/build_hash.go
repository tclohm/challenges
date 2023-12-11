package main

import "fmt"

type Node struct {
	key int
	val int
	next *Node
}

type Hashmap struct {
	container [1000]Node
}

func Init() *Hashmap {
	h := &Hashmap{}
	for i := 0 ; i < 1000 ; i++ {
		h.container[i] = Node{}
	}
	return h
}

func (self *Hashmap) hash(key int) int {
	return key % len(self.container)
}

func (self *Hashmap) put(key int, value int) {
	var current Node = self.container[self.hash(key)]
	for current.next != nil {
		if current.key == key {
			current.val = value
			return
		}
		current = *current.next
	}
	current.next = &Node{key, value, nil}
}

func (self *Hashmap) get(key int) int {
	var current Node = self.container[self.hash(key)]
	for current.next != nil {
		if current.key == key {
			return current.val
		}
		current = *current.next
	}
	return -1
}

func (self *Hashmap) remove(key int) {
	var current Node = self.container[self.hash(key)]
	for current.next != nil {
		if current.next.key == key {
			current.next = current.next.next
			return
		}
		current = *current.next
	}
}

func main() {
	m := Init()
	fmt.Println(m)
}