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
	dummy := &Node{value: 0, next: head}
	p := dummy

	for true {
		kth := getKth(prev, k)
		if kth == nil {
			break
		}

		next := kth.Next

		// reverse
		prev, curr := kth.next, p.next
		for curr != p {
			tmp := curr.next
			curr.next = prev
			prev = curr
			curr = tmp
		}

		tmp := p.next
		p.next = kth
		p = tmp
	}

	return dummy.next
}


func getKth(curr *Node, k int) *Node {
	for curr != nil && k > 0 {
		curr = curr.next
		k -= 1
	}
	return curr
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