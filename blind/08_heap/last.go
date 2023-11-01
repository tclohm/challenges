package main

import (
	"fmt"
	"container/heap"
)

type IntHeap []int
func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, k int) bool { return h[i] < h[k] }
func (h IntHeap) Swap(i, k int) { h[i], h[k] = h[k], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

type KthLargest struct {
	minHeap *IntHeap
	k int
}

func lastStoneWeight(stones []int) int {
	for s := 0 ; s < len(stones) ; s++ {
		stones[s] *= -1
	}

	sh := IntHeap(stones)
	heap.Init(&sh)

	for len(sh) > 1 {
		first, _ := heap.Pop(&sh).(int)
		second, _ := heap.Pop(&sh).(int)
		if second > first {
			heap.Push(&sh, first - second)
		}
	}
	sh = append(sh, 0)
	return abs(stones[0])
}

func abs(x int) int {
	if x < 0 { return -x }
	return x
}

func main() {
	fmt.Println(lastStoneWeight([]int{2,7,4,1,8,1}))
	fmt.Println(lastStoneWeight([]int{1}))
}