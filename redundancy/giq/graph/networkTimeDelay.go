package main

import (
	"fmt"
	"container/heap"
)

type Item struct {
	Name   string
	Expiry int

	Index int // The index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest based on expiration number as the priority
	// The lower the expiry, the higher the priority
	return pq[i].Expiry < pq[j].Expiry
}

// We just implement the pre-defined function in interface of heap.

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// n network nodes labelled 1 to N
// times array, contains edges, [u, v, w]
//						u, source node
//						v, target node
//						w, time taken to travel from source to target
// send signal from k node, return how long
// takes for all nodes to receive the signal
// return -1 if impossible

type doublearray [][]int

func makeArray(n int) []int {
	var distances = make([]int, n, n)
	for i, _ := range distances {
		distances[i] = 100000000
	}
	return distances
}

func timeDelay(times doublearray, n, k int) int {
	var distances = makeArray(n)
	var adjacentList = make(doublearray, n, n)
	distances[k - 1] = 0
	fmt.Println(distances, adjacentList)

	var heap = PriorityQueue{}
	return -1
}

func main() {
	n := 5
	//nodes := []int{1, 2, 3, 4, 5}
	times := [][]int{{1, 2, 9}, {1, 4, 2}, {2, 5, 1}, {4, 2, 4}, {4, 5, 6}, {3, 2, 3}, {5, 3, 7}, {3, 1, 5}}

	timeDelay(times, n, 5)
}