package main

import "fmt"

type Node struct {
	key int
	val int
	next *Node
}

type Hashmap struct {
	array []*Node
	size int
}

func Init() Hashmap {
	return Hashmap{ 
		size: 1000, 
		array: make([]*Node, 1000),
	}
}

func (self *Hashmap) hash(key int) int {
	//fmt.Println("hash", key % len(self.array))
	return key % len(self.array)
}

func (self *Hashmap) put(key int, value int) {
	var position = self.hash(key)
	var newNode *Node = &Node{
		key: key,
		val: value,
		next: nil,
	}

	if self.array[position] == nil {
		self.array[position] = newNode
		return
	}

	var current *Node = self.array[position]

	for current != nil {
		if current.key == key {
			current.val = value
			return
		}
		current = current.next
	}

	newNode.next = self.array[position]
	self.array[position] = newNode
}

func (self *Hashmap) get(key int) int {
	var position = self.hash(key)

	var current *Node = self.array[position]

	for current != nil {
		if current.key == key {
			return current.val
		}
		current = current.next
	}
	return -1
}

func (self *Hashmap) remove(key int) {
	var position = self.hash(key)
	var current *Node = self.array[position]
	
	if current == nil { return }

	if current.key == key {
		current = current.next
		self.array[position] = current
		return
	}

	var prev *Node

	for current != nil {
		if current.key == key {
			prev.next = current.next
			return
		}
		prev = current
		current = current.next
	}
}

func main() {
	m := Init()
	m.put(100, 45)
	m.put(5, 20)
	m.put(999, 1)
	fmt.Println(m.get(4))
	fmt.Println(m.get(5))
}