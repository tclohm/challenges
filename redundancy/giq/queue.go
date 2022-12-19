package main

import "fmt"

type Queue struct {
	Container []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.Container = append(q.Container, item)
	fmt.Println("Enqueue:", item)
}

func (q *Queue) Dequeue() interface{} {
	if q.Empty() {
		return -1
	}
	item := q.Container[0]
	q.Container = q.Container[1:len(q.Container)]
	fmt.Println("Dequeue:", item)
	return item
}

func (q *Queue) Peek() interface{} {
	if q.Empty() {
		return -1
	}
	fmt.Println("Peek:", q.Container[0])
	return q.Container[0]
}

func (q *Queue) Size() int {
	fmt.Println("Queue has", len(q.Container), "items." )
	return len(q.Container)
}

func (q *Queue) Empty() bool {
	if q.Size() > 0 {
		fmt.Println("Queue is not empty")
		return false
	}
	fmt.Println("Queue is empty")
	return true
}

func main() {
	q := Queue{}
	q.Enqueue("Apple")
	q.Enqueue("Banana")
	q.Enqueue("Corn")
	q.Enqueue("Dates")
	q.Peek()
	q.Size()
	q.Dequeue()
	q.Dequeue()
	q.Peek()
	q.Empty()
	q.Dequeue()
	q.Dequeue()
	q.Size()
}