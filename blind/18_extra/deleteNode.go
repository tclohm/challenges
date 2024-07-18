package main

import "fmt"

type Node struct {
	val int
	left *Node
	right *Node
}

func buildTree(nums []int) *Node {
	var root *Node
	var nodes = map[*Node]
	for i := 0 ; i < len(nums) ; i++ {
		node := Node{val: nums[i]}
		if i + 1 < len(nums) {
			node.left = &Node{val: nums[i + 1]}
		}

		if i + 2 < len(nums) {
			node.right = &Node{val: nums[i + 2]}
		}
	}
}
func deleteNodes(root *Node, toDelete []int) []*Node {
	return []int{&Node{val: 0}}
}

func main() {
	fmt.Println(buildTree([]int{1,2,3,4,5,6,7,8}))
}
