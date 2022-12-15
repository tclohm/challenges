package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

func (head *Node) end() *Node {
	for head.Next != nil {
		head = head.Next
	}
	return head
}

func (head *Node) findCycle() Node {
	curr := head

	set := map[int]bool{}

	for curr != nil {
		if _, ok := set[curr.Value]; ok {
			return *curr
		}
		set[curr.Value] = true
		curr = curr.Next
	}
	return *curr
}

func (head *Node) floydCycle() bool {
	var tortoise, hare = head, head

	for hare != nil {
		if tortoise == hare { return true }
		hare = hare.Next.Next
		tortoise = tortoise.Next
	}

	return false
}

func main() {
	n8 := &Node{Value: 8}
	n7 := &Node{Value: 7, Next: n8}
	n6 := &Node{Value: 6, Next: n7}
	n5 := &Node{Value: 5, Next: n6}
	n4 := &Node{Value: 4, Next: n5}
	n3 := &Node{Value: 3, Next: n4}
	n2 := &Node{Value: 2, Next: n3}
	n1 := &Node{Value: 1, Next: n2}

	end := n1.end()
	end.Next = n3

	fmt.Println(n1.findCycle())
	fmt.Println(n1.floydCycle())
}