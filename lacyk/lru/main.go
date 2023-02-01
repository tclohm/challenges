package main

import "fmt"

type Node struct {
	value string
	next *Node
	prev *Node
}

func create(value string) *Node {
	return &Node{value, nil, nil}
}

type LRU struct {
	capacity int
	length int
	head *Node
	tail *Node
	lookup map[string]*Node
	reverseLookup map[*Node]string
}

func Init(capacity int) *LRU {
	l := LRU{}
	l.capacity = capacity
	l.length = 0
	l.lookup = make(map[string]*Node)
	l.reverseLookup = make(map[*Node]string)
	return &l

}

func (self *LRU) update(key string, value string) {
	// does it exist?
	var node, ok = self.lookup[key]
	if !ok {
		node = create(value)
		self.length++
		self.preprend(node)
		self.trimCache()

		self.lookup[key] = node
		self.reverseLookup[node] = key
	} else {
		self.detach(node)
		self.preprend(node)
		node.value = value
	}
	// if it doesn't, insert it
	// - check capacity and evict if over
	// if it does, update to the front of the list
	// and update value
}

func (self *LRU) get(key string) string {
	// check the cache for existence
	var node, ok = self.lookup[key]
	if !ok {
		return "Not here"
	}
	// update the value we found and move it to the front
	self.detach(node)
	self.preprend(node)
	// return out the value found or not found if not exist
	fmt.Println(node.value)
	return node.value
}

func (self *LRU) detach(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	if self.length == 1 {
		self.tail = nil
		self.head = nil
	}

	if self.head == node {
		self.head = self.head.next
	}

	if self.tail == node {
		self.tail = self.tail.prev
	}

	node.next = nil
	node.prev = nil
	return
}

func (self *LRU) preprend(node *Node) {
	if self.head == nil {
		self.head = node
		self.tail = node
		return
	}

	node.next = self.head
	self.head.prev = node
	self.head = node

}

func (self *LRU) trimCache() {
	if self.length <= self.capacity {
		return
	}

	var tail = self.tail
	self.detach(self.tail)

	var key = self.reverseLookup[tail]

	delete(self.lookup, key)
	delete(self.reverseLookup, tail)
	self.length--

}

func main() {
	var lru = Init(3)
	fmt.Println(lru)
	lru.update("foo", "69")
	lru.get("foo")

	lru.update("bar", "420")
	lru.get("foo")

	lru.update("baz", "1337")
	lru.get("baz")

	lru.update("ball", "69420")
	lru.get("ball")
	lru.get("foo")
	lru.get("bar")
	lru.update("foo", "69")
	lru.get("bar")
}