package main

import "fmt"

type Node struct {
	key int
	value int
	next *Node
	prev *Node
}

type DoubleLinkedList struct {
	length int
	head *Node
	tail *Node
}

type NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func (d *DoubleLinkedList) AddLeft(key, value int) {
	node := &Node{
		key: key,
		value: value,
	}
	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		node.next = d.head
		d.head.prev = node
		d.head = node
	}
	d.length++
	return
}

func (d *DoubleLinkedList) RemoveLeft() {
	if d.head == nil {
		return
	} else if d.head == d.tail {
		d.head = nil
		d.tail = nil
	} else {
		d.head = d.head.next
	}
	d.length--
}

func (d *DoubleLinkedList) AddRight(key, value int) {
	node := &Node{
		key: key,
		value: value,
	}
	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		current := d.head
		for current.next != nil {
			current = current.next
		}
		node.prev = current
		current.next = node
		d.tail = node
	}
	d.length++
}

func (d *DoubleLinkedList) Front() *Node {
	return d.head
}

func (d *DoubleLinkedList) MoveToTail(node *Node) {
	prev := node.prev
	next := node.next

	if prev != nil {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}
	if d.tail == node {
		d.tail = prev
	}
	if d.head == node {
		d.head = next
	}
	node.next = nil
	node.prev = nil
	d.length--
	d.AddRight(node)
}

func (d *DoubleLinkedList) Traverse() error {
	if d.head == nil {
		return fmt.Errorf("TraversError: List empty")
	}
	tmp := d.head
	for tmp != nil {
		fmt.Printf("key = %v, value = %v, prev = %v, next = %v\n", tmp.key, tmp.value, tmp.prev, tmp.next)
		tmp = tmp.next
	}
	fmt.Println()
	return nil
}

func (d *DoubleLinkedList) Size() int {
	return d.length
}