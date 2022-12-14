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
	fmt.Println(result + "nil")
}

	// "a" -> "b" -> "c" -> "d" -> "e"
	// 2 - 5
	// "a" -> "e" -> "d" -> "c" -> "b"

	// "a" -> "b" -> "c" -> "d" -> "e"
	// 1 - 3
	// "c" -> "b" -> "a" -> "d" -> "e"

	// "a" -> "b" -> "c" -> "d" -> "e"
	// 3 - 4
	// "a" -> "b" -> "d" -> "c" -> "e"

func (l *ListNode) reverseAt(i, j string) *ListNode {
	// get node
	// store next value
	// update next value
	// store current node
	// update current node to next
	if i == j {
		return l
	}

	var curr = l
	var index = 1

	for curr != nil && curr.Value != i && curr.Next.Value != i {
		curr = curr.Next
		index++
	}

	var leftOut = curr
	curr = curr.Next
	var leftIn = curr
	index++

	for curr != nil && curr.Next.Value != j {
		curr = curr.Next
		index++
	}

	var rightIn = curr
	curr = curr.Next
	var rightOut = curr
	index++

	if index != 1 {
		leftOut.Next = &ListNode{Value: rightIn.Value}
	}
	rightIn.Next = &ListNode{Value: rightIn.Value}

	if index == 1 {
		leftOut.reverse()
		leftOut.Next = rightOut
		return rightIn
	} else {
		leftIn.reverse()
		leftOut.Next = rightIn
		leftIn.Next = rightOut
	}
	return l
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

	n10 := ListNode{Value: "e", Next: &ListNode{}}
	n9  := ListNode{Value: "d", Next: &n10}
	n8  := ListNode{Value: "c", Next: &n9}
	n7  := ListNode{Value: "b", Next: &n8}
	n6  := ListNode{Value: "a", Next: &n7}

	n6.reverseAt("a", "b")

	n6.print()

}