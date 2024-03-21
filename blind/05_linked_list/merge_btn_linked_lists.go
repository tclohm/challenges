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
 
// func mergeInBetween(list1 *Node, a int, b int, list2 *Node) *Node {
    
// }

func main() {
	first := build([]int{10,1,13,6,9,5})
	second := build([]int{1000000,1000001,1000002})

	third := build([]int{0,1,2,3,4,5,6})
	fourth := build([]int{1000000,1000001,1000002,1000003,1000004})

	first.Print()
	second.Print()
	third.Print()
	fourth.Print()
}