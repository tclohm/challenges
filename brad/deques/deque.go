package main

import "fmt"

func reverse(items []interface{}) {
	for i, j := 0, len(items) - 1 ; i < j ; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
}

type Deque struct {
	front_items []interface{}
	rear_items []interface{}
}

func (d *Deque) IsEmpty() bool {
	size := len(d.front_items) + len(d.rear_items)
	return size == 0
}

func (d *Deque) AddFront(item interface{}) {
	d.front_items = append(d.front_items, item)
	reverse(d.front_items)
	deq := append(d.front_items, d.rear_items...)
	fmt.Println("add to front", deq)
}

func (d *Deque) AddRear(item interface{}) {
	d.rear_items = append(d.rear_items, item)
	deq := append(d.front_items, d.rear_items...)
	fmt.Println("add to rear", deq)
}

func (d *Deque) RemoveFront() interface{} {
	item := d.front_items[0]
	d.front_items = d.front_items[1:]
	deq := append(d.front_items, d.rear_items...)
	fmt.Println("remove front", deq)
	return item
}

func (d *Deque) RemoveRear() interface{} {
	item := d.rear_items[0]
	d.rear_items = d.rear_items[1:]
	deq := append(d.front_items, d.rear_items...)
	fmt.Println("remove rear", deq)
	return item
}

func (d *Deque) Size() int {
	return len(d.front_items) + len(d.rear_items)
}

func main() {
	d := Deque{}

	d.AddRear("🧀")
	d.AddFront("🍕")
	d.AddRear("🔌")
	d.AddRear("🔥")
	d.AddFront("🔗")

	fmt.Println(d.Size())

	d.RemoveRear()
	d.RemoveFront()
}