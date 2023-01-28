package main

// heap -- binary tree
// every child and grand child is smaller ( maxheap ) or larger ( minheap ) than the current node
// when node added, we adjust the tree
// when node deleted, we adjust the tree
// the is no traversing the tree

// parent, 
// left child 2(parent position) + 1
// right child 2(parent position) + 2

type MinHeap struct {
	length int
	data []int
}

func (self *MinHeap) insert(value int) {
	self.data[self.length] = value
	self.heapifyUp(self.length)
	self.length++
}

func (self *MinHeap) delete() {
	if self.length == 0 {
		return -1
	}

	if self.length == 1 {
		var out = self.data[0]
		self.data = []
		return out
	}


	self.length--
	self.data[0] = self.data[self.length]
	self.heapifyDown(0)
}

func (self *MinHeap) heapifyDown(idx int) {
	var lIdx = self.leftChild(idx)
	var rIdx = self.rightChild(idx)

	if idx >= self.length || lIdx >= self.length  {
		return
	}

	var lV = self.data[lIdx]
	var rV = self.data[rIdx]
	var v = self.data[idx]

	if lV > rV && v > rV {
		self.data[idx] = rV
		self.data[rIdx] = v
		self.heapifyDown(rIdx)
	} else if rV > lV && v > lV {
		self.data[idx] = lv
		self.data[lIdx] = v
		self.heapifyDown(lIdx)
	}
}

func (self *MinHeap) heapifyUp(idx int) {
	if idx == 0 { return }

	var p = self.parent(idx)
	var parentV = self.data[p]
	var v = self.data[idx]

	if parentV > v {
		self.data[idx] = parentV
		self.data[p] = v
		self.heapifyUp(p)
	}
}

func (self *MinHeap) parent(idx int) int {
	return math.Floor((idx - 1) / 2)
}

func (self *MinHeap) leftChild(idx int) int {
	return idx * 2 + 1
}

func (self *MinHeap) rightChild(idx int) int {
	return idx * 2 + 2
}