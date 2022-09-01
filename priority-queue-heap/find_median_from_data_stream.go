package main

import "container/heap"

type baseheap []int
type minheap struct { baseheap }
type maxheap struct { baseheap }

type MedianFinder struct {
	Min *minheap
	Max *maxheap
}

// boilerplate code to satisfy heap.Interface
func (self baseheap) Len() int { 
	return len(self) 
}

func (self baseheap) Peek() int { 
	return self[0] 
}

func (self baseheap) Swap(a, b int) { 
	self[a], self[b] = self[b], self[a] 
}

func (self *baseheap) Push(x interface{}) { 
	*self = append(*self, x.(int)) 
}

func (self *baseheap) Pop() interface{} {
	x := (*self)[len(*self) - 1]
	*self = (*self)[:len(*self) - 1]
	return x
}

func (self minheap) Less(a, b int) bool {
	return self.baseheap[a] < self.baseheap[b]
}

func (self maxheap) Less(a, b int) bool {
	return self.baseheap[a] > self.baseheap[b]
}	

func Constructor() MedianFinder {
	return MedianFinder{new(minheap), new(maxheap)}
}

// track numbers in two heaps, Min and Max
// the value of num will deterMine which heap to push into
// after pushing the length of the heap will be balanced to make
// sure they are within 1 length of eachother
func (self *MedianFinder) AddNum(num int) {
	if self.Min.Len() > 0 && num <= self.Min.Peek() {
		heap.Push(self.Min, num)
		if self.Min.Len() > self.Max.Len() + 1 {
			heap.Push(self.Max, heap.Pop(self.Min))
		}
	} else {
		heap.Push(self.Max, num)
		if self.Max.Len() > self.Min.Len() + 1 {
			heap.Push(self.Min, heap.Pop(self.Max))
		}
	}
}

// return the median value of all the numbers added in AddNum
func (self *MedianFinder) FindMedian() float64 {
	if self.Min.Len() > self.Max.Len() {
		return float64(self.Min.Peek())
	} else if self.Min.Len() < self.Max.Len() {
		return float64(self.Max.Peek())
	}
	return float64(self.Min.Peek() + self.Max.Peek()) / 2.0
}