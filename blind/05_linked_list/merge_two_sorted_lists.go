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

func build(nums []int) *Node {
	var prev = &Node{value: nums[0]}
	var head = prev
	for i := 1 ; i < len(nums) ; i++ {
		n := &Node{value: nums[i]}
		prev.next = n
		prev = n
	}

	return head
}

func merge(list1 *Node, list2 *Node) *Node {
	node1, node2 := list1, list2

	if node1 == nil { return node2 }
	if node2 == nil { return node1 }

	var current *Node = &Node{}
	var head = current

	for node1 != nil || node2 != nil {
		if node1.value <= node2.value {
			current.value = node1.value
			node1 = node1.next
		} else {
			current.value = node2.value
			node2 = node2.next
		}
		current.next = &Node{}
		current = current.next

		if node1 == nil || node2 == nil { break }
	}

	for node1 == nil && node2 != nil {
		current.value = node2.value
		current.next = &Node{}
		current = current.next
		node2 = node2.next
	}

	for node2 == nil && node1 != nil {
		current.value = node1.value
		current.next = &Node{}
		current = current.next
		node1 = node1.next
	}

	return head
}

func main() {
	odd := build([]int{1, 3, 5, 7, 9})
	even := build([]int{2, 4, 6, 8})
	odd.Print()
	even.Print()

	merged := merge(even, odd)
	merged.Print()
}