package main

import "fmt"

type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	head *Node
}

func Constructor(nums []int) LinkedList {
	list := LinkedList{
		head: &Node{value: nums[0]},
	}
	current := list.head
	for i := 1 ; i < len(nums) ; i++ {
		current.next = &Node{value: nums[i]}
		current = current.next
	}

	return list
}

func (self *LinkedList) findMiddle() int {
	slow, fast := self.head, self.head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow.value
}

func main() {
	linked := Constructor([]int{1,2,3,4,5,6,7,8})
	fmt.Println(linked.findMiddle())
}