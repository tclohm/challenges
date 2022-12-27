package main

import (
	"fmt"
	"math"
)

// heaps - complete binary tree
// max heap - values with children are smaller than it and below, root node has the greatest value
// min heap - values of children are greater than it, root node is smallest value

// [50, 40, 25, 20, 35, 10, 15] 
//  0   1   2   3   4   5   6
// parent : floor((idx - 1) / 2)
// left   : (idx * 2) + 1
// right  : (idx * 2) + 2


type PriorityQueue struct {
	Heap []int{}
	Comparator func(a, b int) bool { return a > b }
}

func (pq *PriorityQueue) Size() int {
	return len(pq.Heap)
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Size() == 0
}

func (pq *PriorityQueue) Peek() {
	return pq.Heap[0]
}

func (pq *PriorityQueue) parent(index int) int {
	return math.Floor((index - 1) / 2)
}

func (pq *PriorityQueue) leftChild(index int) int {
	return index * 2 + 1
}

func (pq *PriorityQueue) rightChild(index int) int {
	return index * 2 + 2
}

func (pq *PriorityQueue) swap(i, j int) {
	pq.Heap[i], pq.Heap[j] = pq.Heap[j], pq.Heap[i]
}

func (pq *PriorityQueue) compare(a, b int) bool {
	return pq.Comparator(a, b)
}

func (pq *PriorityQueue) Push(value int) int {
	pq.Heap = append(pq.Heap, value)
	pq.siftUp()
	return pq.Size()
}

func (pq *PriorityQueue) siftUp() {
	var nodeIndex = pq.Size() - 1
	for nodeIndex > 0 && pq.compare(nodeIndex, pq.parent(nodeIndex)) {
		pq.swap(nodeIndex, pq.parent(nodeIndex))
		nodeIndex = pq.parent(nodeIndex)
	}
}

func (pq *PriorityQueue) siftDown() {
	var nodeIndex = 0
	for pq.leftChild(nodeIndex) < pq.Size() && 
	pq.compare(pq.leftChild(nodeIndex), nodeIndex) || 
	pq.rightChild(nodeIndex) < pq.Size() && 
	pq.compare(pq.rightChild(nodeIndex), nodeIndex) {
		var greaterNodeIndex
		if pq.rightChild(nodeIndex) < pq.Size() &&
		pq.compare(pq.rightChild(nodeIndex), pq.leftChild(nodeIndex)) {
			greaterNodeIndex = pq.rightChild(nodeIndex)
		} else {
			greaterNodeIndex = pq.leftChild(nodeIndex)
		}

		pq.swap(greaterNodeIndex, nodeIndex)
		nodeIndex = greaterNodeIndex
	}
}

func (pq *PriorityQueue) Pop() {
	if pq.Size() > 1 {
		pq.swap(0, pq.Size() - 1)
	}
	var popped = pq.Heap()
	pq.siftDown()
	return popped
}