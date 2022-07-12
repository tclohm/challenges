package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func reorder(head *ListNode) {
	if head == nil || head.Next == nil { return }
	var prev *ListNode = nil
	var slow *ListNode = head
	var fast *ListNode = head
	var l1   *ListNode = head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = nil
	l2 := reverse(slow)
	merge(l1, l2)
}

func reverse(node *ListNode) *ListNode {
	var prev *ListNode = nil
	var curr *ListNode = node

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func merge(l1, l2 *ListNode) {
	for l1 != nil {
		next1, next2 := l1.Next, l2.Next
		l1.Next = l2
		if next1 == nil { break }
		l2.Next = next1
		l1 = next1
		l2 = next2
	}
}

func __print__(current *ListNode) string {
	result := ""
	for current != nil {
		if current.Next != nil {
			result +=  fmt.Sprintf("%v -> ", current.Val)
		} else {
			result += fmt.Sprintf("%v", current.Val)
		}
		
		current = current.Next
	}
	return result
}


func main() {
	five := &ListNode{Val: 5, Next: nil}
	four := &ListNode{Val: 4, Next: five}
	three := &ListNode{Val: 3, Next: four}
	two := &ListNode{Val: 2, Next: three}
	one := &ListNode{Val: 1, Next: two}

	fmt.Println(__print__(one))
	reorder(one)
	fmt.Println(__print__(one))
}