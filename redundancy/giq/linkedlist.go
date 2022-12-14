package main

import (
	"fmt"
)

type ListNode struct {
	Value string
	Next *ListNode
}

type ListNodeInt struct {
	Value int
	Next *ListNodeInt
}


func (l *ListNode) reverse() *ListNode {
	var curr = l
	var prev = &ListNode{}

	for curr.Next != nil {
		var next *ListNode = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func (l *ListNode) print() {
	var result = ""
	var curr = l
	for curr.Next != nil {
		result += curr.Value + " -> "
		curr = curr.Next
	}
	fmt.Println(result + "nil")
}

func (l *ListNodeInt) print() {
	var result = ""
	var curr = l
	for curr.Next != nil {
		result += fmt.Sprint(curr.Value, " -> ")
		curr = curr.Next
	}
	fmt.Println(result + "nil")
}

	// 1 -> 2 -> 3 -> 4 -> 5
	// 2 - 5
	// 1 -> 5 -> 4 -> 3 -> 2

	// 1 -> 2 -> 3 -> 4 -> 5
	// 1 - 3
	// 3 -> 2 -> 1 -> 4 -> 5

	// 1 -> 2 -> 3 -> 4 -> 5
	// 3 - 4
	// 1 -> 2 -> 4 -> 3 -> 5

func (l *ListNodeInt) reverseAt(j, k int) *ListNodeInt {
	
}

func main() {
	n5 := ListNode{Value: "Taylor", Next: &ListNode{}}
	n4 := ListNode{Value: "I'm", Next: &n5}
	n3 := ListNode{Value: ",", Next: &n4}
	n2 := ListNode{Value: "World", Next: &n3}
	n1 := ListNode{Value: "Hello", Next: &n2}

	var curr = &n1

	curr.print()

	fmt.Println("\nreversing!\n")

	curr = &n1
	
	var head = curr.reverse()

	head.print()

	fmt.Println("\nreverse at\n")

	n10 := ListNodeInt{Value: 5, Next: &ListNodeInt{}}
	n9  := ListNodeInt{Value: 4, Next: &n10}
	n8  := ListNodeInt{Value: 3, Next: &n9}
	n7  := ListNodeInt{Value: 2, Next: &n8}
	n6  := ListNodeInt{Value: 1, Next: &n7}

	n6.reverseAt(1, 2)

	n6.print()

}