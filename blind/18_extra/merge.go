package main

import (
	"fmt"
	"bytes"
)

type Node struct {
	val int
	next *Node
}

func buildLinkedList(nums []int) *Node {
	var tail *Node = &Node{}
	var head *Node = &Node{nums[0], tail}		
	for i := 1 ; i < len(nums) ; i++ {
		tail.val = nums[i]
		if i < len(nums) - 1 {
			tail.next = &Node{0, nil}
			tail = tail.next
		}
	}
	
	return head
}

func printLinkedList(head *Node) {
	var buffer bytes.Buffer
	for head != nil {
		buffer.WriteString(fmt.Sprintf("%d -> ", head.val))
		head = head.next
	}
	buffer.WriteString("nil")

	fmt.Println(buffer.String())
}

func isZero(head *Node) bool {
	if head.val == 0 {
		return true
	}
	return false
}

func mergeNodesBetweenZero(head *Node) []int {
	var result = []int{}
	var add = 0
	for nil != head {
		if head.val != 0 {
			add += head.val
		}
		if head.val == 0 && add != 0 {
			result = append(result, add)
			add = 0
		}
		head = head.next
	}
	return result
}

func main() {
	one := buildLinkedList([]int{0,3,1,0,4,5,2,0})
	printLinkedList(one)
	fmt.Println(mergeNodesBetweenZero(one))
	two := buildLinkedList([]int{0,1,0,3,0,2,2,0})
	printLinkedList(two)
	fmt.Println(mergeNodesBetweenZero(two))
}
