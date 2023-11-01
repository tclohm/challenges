package main

import "fmt"

type Node struct {
	value int
	next *Node
}

func (n *Node) Print() {
	for n != nil {
		if n.next == nil {
			fmt.Printf("%d\n", n.value)
			return
		} 
		fmt.Printf("%d -> ", n.value)
		n = n.next
	}
}

func (n *Node) Reverse() *Node {
	var prev  *Node
	curr := n
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	return prev
}

func reverseKGroup(head *Node, k int) *Node {
	return &Node{}
}


func main() {
	Node5 := Node{value: 5, next: nil}
	Node4 := Node{value: 4, next: &Node5}
	Node3 := Node{value: 3, next: &Node4}
	Node2 := Node{value: 2, next: &Node3}
	Node1 := Node{value: 1, next: &Node2}
	Node1.Print()
	Node1.Reverse()
	fmt.Println("\nReversed")
	Node5.Print()
	fmt.Println("\nReversed again")
	Node5.Reverse()
	Node1.Print()
}