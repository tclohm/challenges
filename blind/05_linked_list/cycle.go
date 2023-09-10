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

func (n *Node) hasCycle() int {
	var ht = make(map[int]int)
	var index = 0
	for n != nil {
		if _, exist := ht[n.value]; exist {
			return ht[n.value]
		} 
		ht[n.value] = index
		index++
		n = n.next
	}

	return -1
}

func main() {
	basic := build([]int{3,2,0,-4, 2})
	one := build([]int{1,2,1})
	two := build([]int{1})

	// basic.Print()
	// one.Print()
	// two.Print()

	fmt.Println(basic.hasCycle())
	fmt.Println(one.hasCycle())
	fmt.Println(two.hasCycle())

}