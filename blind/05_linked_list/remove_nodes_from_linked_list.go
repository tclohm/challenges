package main

import (
	"fmt"
)

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

// remove every node which has a node with a greater value anywhere to the right side of it
// present node : 1
// next node : 2
// question is our present node greater than next node
// if so keep it
// if not remove it
// 1, 2, 4, 19, 4, 3, 2
func removeNodes(head *Node) *Node {
	// is the current node larger than the next
	// if it is add them together
	// move to the next node
	if head.next == nil {
		return head
	}

	head.next = removeNodes(head.next)
	if head.value < head.next.value {
		return head.next
	}

	return head
}


func reverse(head *Node) *Node {
	var prev, curr *Node = nil, head
	for curr != nil {
		// we need next to continue
		next := curr.next
		// switch currents next to the previous value
		curr.next = prev
		// update our previous and current 
		prev = curr
		curr = next
	}
	return prev
}

// linear time solution
func removeNodesWithStack(head *Node) *Node {
	stack := []int{}
	current := head
	for current != nil {
		for len(stack) > 0 && current.value > stack[len(stack) - 1] {
			stack = stack[:len(stack) - 1]
		}
		stack = append(stack, current.value)
		current = current.next
	}

	dummy := &Node{}
	current = dummy
	for _, n := range stack {
		current.next = &Node{value: n}
		current = current.next
	}
	return dummy.next
}

func main() {
	b1 := build([]int{5,2,13,3,8})
	b2 := build([]int{1,1,1,1})

	b1.Print()
	b2.Print()

	newB1 := removeNodesWithStack(b1)
	newB2 := removeNodesWithStack(b2)

	newB1.Print()
	newB2.Print()
}