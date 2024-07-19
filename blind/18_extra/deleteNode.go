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

func printTree(root *Node, nodes map[int]bool) map[int]bool {
	if root != nil && !nodes[root.val] {
			nodes[root.val] = true
	}

	if nil != root.left && nodes[root.val] {
		printTree(root.left, nodes)
	}

	if nil != root.right && nodes[root.val]  {
		printTree(root.right, nodes)
	}
	
	return nodes
}

func deleteNodes(root *Node, toDelete []int) []*Node {
	return []*Node{&Node{val: 0}}
}

func main() {
	r := buildTree([]int{1,2,3,4,5,6,7,8})
	fmt.Println(printTree(r, map[int]bool{}))
}
