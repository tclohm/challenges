package main

import (
	"fmt"
	"container/heap"
)

// one ladder or difference (height[i + 1] - height[i]) in bricks
func furthestBuildingNOPE(heights []int, bricks int, ladders int) int {
	// use bricks first
	// or if the bricks dont make the difference to reach next building use ladder

	for i := 0 ; i < len(heights) - 1 ; i++ {
			height := heights[i]
			next_height := heights[i + 1]
			if height > next_height { 
				continue 
			}
			
			if bricks == 0 && ladders == 0 { 
				return i
			}
			
			if height < next_height && (next_height - height) <= bricks {
				bricks -= (next_height - height)
				continue
			} else {
				if ladders > 0 {
					ladders--
				}
			}

	}
	return len(heights) - 1
}

type IntHeap []int

func (self IntHeap) Less(i, j int) bool {
	return self[i] < self[j]
}

func (self IntHeap) Swap(i, j int) {
	self[i], self[j] = self[j], self[i] 
}

func (self IntHeap) Len() int {
	return len(self)
}

func (self *IntHeap) Pop() interface{} {
	old := *self
	length := len(*self)
	last := old[length - 1]
	*self = old[:length - 1]

	return last
}

func (self *IntHeap) Push(integer interface{}) {
	*self = append(*self, integer.(int))
}


func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := &IntHeap{}
	heap.Init(h)

	for i := 1 ; i < len(heights) ; i++ {
		difference := heights[i] - heights[i - 1]
		if difference > 0 {
			heap.Push(h, difference)
			if h.Len() > ladders { bricks -= heap.Pop(h).(int) }
			if bricks < 0 { return i - 1 }
		}
	}
	return len(heights) - 1
}

func main() {
	fmt.Println(furthestBuilding([]int{4,2,7,6,9,14,12}, 5, 1))
	fmt.Println(furthestBuilding([]int{4,12,2,7,3,18,20,3,19}, 10, 2))
	fmt.Println(furthestBuilding([]int{14,3,19,3}, 17, 0))
}