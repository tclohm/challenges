package main

import "fmt"

type Node struct {
	val int
	left *Node
	right *Node
}

func buildTree(nums []int) *Node {
	for i, n := range nums {
		node := Node{}
	}
}

func deleteNodes(root *Node, toDelete []int) []*Node {
	return []int{&Node{val: 0}}
}

func main() {
	fmt.Println(deleteNodes())
}
