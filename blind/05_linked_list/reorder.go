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

func reorder(head *Node) {
	slow := head
	fast := head.next

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}

	reversed := reverse(slow.next)
	slow.next = nil

	current := head

	for current != nil && reversed != nil {
		next := current.next
		reversedNext := reversed.next
		current.next = reversed
		reversed.next = next
		current = next
		reversed = reversedNext
	}
}

func reverse(node *Node) *Node {
	var prev, current *Node = nil, node

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

func main() {
	odd := build([]int{1, 3, 5, 7, 9})
	even := build([]int{2, 4, 6, 8})
	odd.Print()
	even.Print()

	reorder(odd)
	reorder(even)
	odd.Print()
	even.Print()
}