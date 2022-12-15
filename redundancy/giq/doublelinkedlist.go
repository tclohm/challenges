package main

import (
	"fmt"
	//"reflect"
)

type ListNode struct {
	Value interface{}
	Prev *ListNode
	Next *ListNode
	Child *ListNode
}

func build(nodes []interface{}) *ListNode {
	var head = &ListNode{Value: nodes[0]}
	var current = head
	fmt.Println("nodes", nodes[0])
	for i := 1 ; i < len(nodes) ; i++ {
		var val = nodes[i]
		var next *ListNode
		switch val.(type) {
		default:
			next = build(val)
			current.Child = next
			continue
		case int:
			next = &ListNode{Value: val}
			current.Next = next
			next.Prev = current
			current = next
		} 
	}

	return head
}
func main() {
	// MARK: -- crazy
	var nodes = []interface{}{1, []interface{}{2, 7, []int{8, 10, 11,}, 9,}, 3, 4, []int{5, 12, 13,}, 6,}
	fmt.Println(nodes)
	build(nodes)

}