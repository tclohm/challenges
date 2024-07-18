package main

import "fmt"

type Node struct {
	val int
	left *Node
	right *Node
}

func buildTree(nums []int) *Node {
	var root *Node
	root = &Node{val: nums[0]}
	if 1 < len(nums) {
		left := buildTree(nums[1:])
		root.left = left
	}
	if 2 < len(nums) {
		right := buildTree(nums[2:])
		root.right = right
	}
	return root
}

func printTree(root *Node, nodes *[]int) []int {
	if root != nil && !in(root, *nodes) {
		*nodes = append(*nodes, root.val)
	}

	if nil != root.left && !in(root.left, *nodes) {
		printTree(root.left, nodes)
	}

	if nil != root.right && !in(root.right, *nodes)  {
		printTree(root.right, nodes)
	}

	return *nodes
}

func in(n *Node, visited []int) bool {
	for number := range visited {
		if n.val == number {
			return true
		}
	}
	return false
}


func deleteNodes(root *Node, toDelete []int) []*Node {
	return []*Node{&Node{val: 0}}
}

func main() {
	r := buildTree([]int{1,2,3,4,5,6,7,8})
	fmt.Println(printTree(r, &[]int{}))
}
