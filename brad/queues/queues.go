package main

import "fmt"

type Queue struct {
	_items []interface{}
}

func (q *Queue) IsEmpty() bool {
	return len(q._items) == 0
}

func (q *Queue) Enqueue(item interface{}) {
	fmt.Println("enqueue", item)
	q._items = append(q._items, item) 
	fmt.Println("after enqueued:", q._items)
}

func (q *Queue) Dequeue() {
	fmt.Println("dequeue")
	q._items = q._items[1:]
	fmt.Println("after dequeue:", q._items)
}

func (q *Queue) Size() int {
	return len(q._items)
}

func main() {
	q := Queue{}
	q.Enqueue("Hello")
	q.Enqueue(9)
	q.Enqueue(1.0912)
	q.Enqueue("🍕")
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
}