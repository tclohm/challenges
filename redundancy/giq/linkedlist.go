package main

import "fmt"

type ListNode struct {
	Value string
	Next *ListNode
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
	fmt.Println(result + "Nil")
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
}