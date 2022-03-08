package main

import "fmt"

type Node struct {
	value interface{}
	next *Node
}

type UnorderedList struct {
	head *Node
}

func (ul *UnorderedList) IsEmpty() bool {
	return ul.head == nil
}

func (ul *UnorderedList) Add(item interface{}) {
	fmt.Println("adding", item)
	temp := &Node{value: item}
	temp.next = ul.head
	ul.head = temp
}

func (ul *UnorderedList) Size() int {
	var current *Node = ul.head
	var count int = 0

	for current != nil {
		count++
		current = current.next
	}

	return count
}

func (ul *UnorderedList) Search(item interface{}) bool {
	var current *Node = ul.head

	for current != nil {
		if current.value == item {
			return true
		}
		current = current.next
	}
	return false
}

func (ul *UnorderedList) Remove(item interface{}) {
	var current *Node = ul.head
	var previous *Node

	for {
		if current.value == item {
			break
		}
		previous, current = current, current.next
	}

	if previous == nil {
		ul.head = current.next
	} else {
		previous.next = current.next
	}
}

func main() {
	ul := UnorderedList{}
	n := Node{value: 93}
	fmt.Println(ul, n.value)
	fmt.Println(ul.IsEmpty())
	ul.Add("hello world")
	ul.Add(13)
	ul.Add("🍕")
	ul.Add("T Money")

	fmt.Println("Size", ul.Size())
	fmt.Println("👻 is in the list?", ul.Search("👻"))

	ul.Add("👻")
	ul.Add("👾")
	ul.Add("The end is near")
	ul.Add("🤖")

	fmt.Println("Size", ul.Size())

	fmt.Println("🍕 is in the list?", ul.Search("🍕"))

	fmt.Println("Removing 🤖")
	ul.Remove("🤖")

	fmt.Println("Is 🤖 in the list?", ul.Search("🤖"))

	ul.Remove("hello world")
}