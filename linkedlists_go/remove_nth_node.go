package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
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

func remove(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 { return head }
	preHeader := &ListNode{Next: head}
	left, right := preHeader, head

	for i := 0 ; right != nil && i < n ; i++ {
		right = right.Next
	}

	for right != nil {
		left, right = left.Next, right.Next
	}

	left.Next = left.Next.Next

	return preHeader.Next
}

func main() {
	eight := &ListNode{Val: 8, Next: nil}
	seven := &ListNode{Val: 7, Next: eight}
	six := &ListNode{Val: 6, Next: seven}
	five := &ListNode{Val: 5, Next: six}
	four := &ListNode{Val: 4, Next: five}
	three := &ListNode{Val: 3, Next: four}
	two := &ListNode{Val: 2, Next: three}
	one := &ListNode{Val: 1, Next: two}

	fmt.Println(__print__(one))

	remove(one, 5)

	fmt.Println(__print__(one))
}