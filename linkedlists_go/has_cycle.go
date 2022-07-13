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

func hasCycle(head *ListNode) bool {
	hm := make(map[*ListNode]bool)
	for head != nil {
		if _, ok := hm[head]; ok {
			return true
		} else {
			hm[head] = true
		}
		head = head.Next
	}
	return false
}

func hasCyclePointers(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { return true }
	}
	return false
}

func main() {
	var nFour, zero, two, otherTwo, three *ListNode
	otherTwo = &ListNode{Val: 2, Next: three}
	nFour = &ListNode{Val: -4, Next: otherTwo}
	zero = &ListNode{Val: 0, Next: nFour}
	two = &ListNode{Val: 2, Next: zero}
	three = &ListNode{Val: 3, Next: two}

	fmt.Println(hasCycle(three))
}