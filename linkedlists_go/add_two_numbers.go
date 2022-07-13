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

func exist(node *ListNode) bool {
	if node != nil && node.Val > 0 {
		return true
	}
	return false
}

func addTwoNumber(l1, l2 *ListNode) *ListNode {
	carry := 0
	head := &ListNode{}
	current := head
	for l1 != nil || l2 != nil {
		value1 := 0
		value2 := 0

		if exist(l1) {
			value1 = l1.Val
		} else {
			value1 = 0
		}

		if exist(l2) {
			value2 = l2.Val
		} else {
			value2 = 0
		}

		sum := value1 + value2 + carry
		carry = sum / 10

		new := &ListNode{Val: sum % 10}

		current.Next = new
		current = new

		l1 = l1.Next
		l2 = l2.Next
	}

	if carry > 0 {
		new := &ListNode{Val: carry}
		current.Next = new
		current = new
	}

	return head.Next
}

func main() {
	otherFour := &ListNode{Val: 4, Next: nil}
	six := &ListNode{Val: 6, Next: otherFour}
	five := &ListNode{Val: 5, Next: six}

	three := &ListNode{Val: 3, Next: nil}
	four := &ListNode{Val: 4, Next: three}
	two := &ListNode{Val: 2, Next: four}

	
	fmt.Println(__print__(two))
	fmt.Println(__print__(five))
	fmt.Println("------------")
	new := addTwoNumber(two, five)
	fmt.Println(__print__(new))
	fmt.Println()
	zero := &ListNode{Val: 0, Next: nil}
	otherZero := &ListNode{Val: 0, Next: nil}
	fmt.Println(__print__(zero))
	fmt.Println(__print__(otherZero))
	zeros := addTwoNumber(zero, otherZero)
	fmt.Println("------------")
	fmt.Println(__print__(zeros))

	nineThree := &ListNode{Val: 9, Next: nil}
	nineTwo := &ListNode{Val: 9, Next: nineThree}
	nineOne := &ListNode{Val: 9, Next: nineTwo}

	otherNineThree := &ListNode{Val: 9, Next: nil}
	otherNineTwo := &ListNode{Val: 9, Next: otherNineThree}
	otherNine := &ListNode{Val: 9, Next: otherNineTwo}
	fmt.Println(__print__(nineOne))
	fmt.Println(__print__(otherNine))

	nines := addTwoNumber(nineOne, otherNine)
	fmt.Println(__print__(nines))
}