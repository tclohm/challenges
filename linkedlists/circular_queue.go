package main

import "fmt"

type CircularQueue struct {
	size int
	nodes []interface{}
	head int
	last int
}

func NewQueue(num int) *CircularQueue {
	var circleQueue CircularQueue
	circleQueue = CircularQueue{size: num + 1, head: 0, last: 0}
	circleQueue.nodes = make([]interface{}, circleQueue.size)
	return &circleQueue
}

func (circularQueue CircularQueue) IsUnUsed() bool {
	return circularQueue.head == circularQueue.last
}

func (circularQueue CircularQueue) IsComplete() bool {
	return circularQueue.head == (circularQueue.last + 1) % circularQueue.size
}

func (circularQueue *CircularQueue) Add(element interface{}) {
	if circularQueue.IsComplete() {
		panic("Queue is Completely Utilized")
	}
	circularQueue.nodes[circularQueue.last] = element
	circularQueue.last = (circularQueue.last + 1) % circularQueue.size
}

func (circularQueue *CircularQueue) MoveOneStep() (element interface{}) {
	if circularQueue.IsUnUsed() {
		return nil
	}

	element = circularQueue.nodes[circularQueue.head]
	circularQueue.head = (circularQueue.head + 1) % circularQueue.size
	return
}

func main() {
	var cq *CircularQueue
	cq = NewQueue(5)

	cq.Add(1)
	cq.Add(2)
	cq.Add(3)
	cq.Add(4)
	cq.Add(5)

	fmt.Println(cq.nodes)
}