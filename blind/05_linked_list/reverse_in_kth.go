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

func reverseKGroup(head *Node, k int) *Node {
	dummy := &Node{value: 0, next: head}
	groupPrev := dummy

	for true {
		kth := getKth(groupPrev, k)
		if kth == nil {
			break
		}

		groupNext := kth.next

		// reverse
		prev, curr := kth.next, groupPrev.next
		for curr != groupNext {
			tmp := curr.next
			curr.next = prev
			prev = curr
			curr = tmp
		}

		tmp := groupPrev.next
		groupPrev.next = kth
		groupPrev = tmp
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

	reverseKGroup(&Node1, 2).Print()
}