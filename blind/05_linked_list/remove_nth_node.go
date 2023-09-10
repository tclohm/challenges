package main

import "fmt"

type Node struct {
	value int
	next *Node
}

func (n *Node) Print() {
	for n != nil {
		if n.next == nil {
			fmt.Printf("%d\n", n.value)
			return
		} 
		fmt.Printf("%d -> ", n.value)
		n = n.next
	}
}

func build(nums []int) *Node {
	var prev = &Node{value: nums[0]}
	var head = prev
	for i := 1 ; i < len(nums) ; i++ {
		n := &Node{value: nums[i]}
		prev.next = n
		prev = n
	}

	return head
}

func removeNthFromEnd(head *Node, n int) *Node {
	if head == nil { return head }

	slow, fast := head, head

	for n > 0 {
		fast = fast.next
		n--
	}

	if fast == nil { return slow.next }

	if fast.next != nil {
		slow = slow.next
		fast = fast.next
	}

	slow.next = slow.next.next
	return head
}

func main() {
	basic := build([]int{1, 2, 3, 4, 5})
	one := build([]int{1})
	two := build([]int{1,2})

	basic.Print()
	one.Print()
	two.Print()

	removeNthFromEnd(basic,2).Print()
	removeNthFromEnd(one,1).Print()
	removeNthFromEnd(two,1).Print()

}