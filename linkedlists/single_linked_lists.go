package main

import "fmt"

type Node struct {
	next *Node
	value rune
}

func Create() *Node {
	head := &Node{nil, 'a'}

	curr := head

	for i := 'b' ; i <= 'z' ; i++ {
		node := &Node{nil, i}
		curr.next = node
		curr = node
	}

	return head
}

func Reverse(list *Node) *Node {
	var curr *Node
	var top *Node

	curr = list
	top = nil

	for {
		if curr == nil {
			break
		}

		tmp := curr.next
		curr.next = top
		top = curr
		curr = tmp
	}

	return top
}

func StringifyList(list *Node) {
	curr := list
	output := ""
	for curr != nil {
		output += string(curr.value) + " -> "
		curr = curr.next
	}
	fmt.Println(output + "nil")
}

func main() {
	var linkedlist = Create()
	StringifyList(linkedlist)
	StringifyList(Reverse(linkedlist))
}