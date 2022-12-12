package main

import "fmt"

type ListNode struct {
	Value string
	Next *ListNode
}

func main() {
	n := ListNode{Value: "World", Next: &ListNode{}}
	l := ListNode{Value: "Hello", Next: &n}

	fmt.Println(l.Value)
}