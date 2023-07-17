package main

type Node struct {
	key int
	val int
	prev *Node
	next *Node
}

type LRUCache struct {
	capacity int
	cache map[int]*Node
	left *Node
	right *Node
}

func New(capacity int) LRUCache {
	ret := LRUCache{
		capacity: 	capacity,
		cache: 		make(map[int]*Node),
		left : 		&Node{key: 0, val: 0},
		right: 		&Node{key: 0, val: 0}
	}
	ret.left.next, ret.right.prev = ret.right, ret.left
	return ret
}

func (self *LRUCache) Remove(node *Node) {
	prev, next := node.prev, node.next
	prev.next, next.prev = next, prev
}

func (self *LRUCache) Insert(node *Node) {
	prev, next := self.right.prev, self.right
	next.prev = node
	prev.next = next.prev
	node.next, node.prev = next, prev
}

func (self *LRUCache) Get(key int) int {
	if _, ok := self.cache[key]; ok {
		self.Remove(self.cache[key])
		self.Insert(self.cache[key])
		return self.cache[key].val
	}
	return -1
}

func (self *LRUCache) Put(key int, value int) {
	if _, ok := self.cache[key]; ok {
		self.Remove(self.cache[key])
	}
	self.cache[key] = &Node{key: key, val: value}
	self.Insert(self.cache[key])

	if len(self.cache) > self.capacity {
		lru := self.left.next
		self.Remove(lru)
		delete(self.cache, lru.key)
	}
}