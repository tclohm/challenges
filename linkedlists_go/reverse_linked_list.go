package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		// grab the next node and then point the current node's next to prev 		
		next := current.Next
		current.Next = prev
		// move previous to current and current to next
		prev = current
		current = next
	}

	return prev
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
	fmt.Println("reverse")
	reverse(one)
	fmt.Println("reversed")
	fmt.Println(__print__(five))
}