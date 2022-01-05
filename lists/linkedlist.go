package main

import (
	"fmt"
	_"log"
	"strconv"
)

type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (linkedList *LinkedList) Last() *Node {
	var node *Node
	var last *Node

	for node = linkedList.head ; node != nil ; node = node.next {
		if node.next == nil {
			last = node
		}
	}

	return last
}

func (linkedList *LinkedList) Shift(value int) {
	var node = Node{}

	node.value = value

	if node.next == nil {
		node.next = linkedList.head
	}
	linkedList.head = &node
}

func (linkedList *LinkedList) Push(value int) {
	var node = &Node{}

	node.value = value
	node.next = nil

	tail := linkedList.Last()

	if tail != nil {
		tail.next = node
	}

}

func (linkedList *LinkedList) Find(value int) *Node {
	var node *Node
	var target *Node
	for node = linkedList.head ; node != nil ; node = node.next {
		if node.value == value {
			target = node
		}
	}
	return target
}

func (linkedList *LinkedList) PushAfter(nodeValue int, value int) {
	var node = &Node{}
	node.value = value
	node.next = nil

	nodeWith := linkedList.Find(nodeValue)
	if nodeWith != nil {
		// node next is nodeWith next
		// then nodeWith next is node
		node.next = nodeWith.next
		nodeWith.next = node
	}
}

func (linkedList *LinkedList) Print() {
	var printout = ""

	var current = linkedList.head

	for current != nil {
		string := strconv.Itoa(current.value)
		printout += string + " -> "
		current = current.next
	}
	printout += "nil"

	fmt.Println(printout)
}



func main() {
	var linkedList LinkedList
	linkedList.Shift(1)
	linkedList.Shift(3)
	linkedList.Shift(42)
	linkedList.Shift(69)
	linkedList.Shift(0)

	linkedList.Push(420)

	linkedList.PushAfter(3, 99)

	linkedList.Print()
}