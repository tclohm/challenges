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

func deleteNodes(root *Node, toDelete []int) [][]*Node {
	// delete set
	var deleting = map[int]bool{}
	// result set
	var result = map[[]int]bool{}

	// dfs func
	var dfs func(node *Node)
	dfs = func(node *Node) {
	// if not node return none
		if nil == node { return }
		var res = node
	// if node.val in delete set
		if deleting[node.val] {
			res = nil
			delete(result, node)
			if node.left != nil {
				result[node.left] = true
			}
			if node.right != nil {
				result[node.right] = true
			}
		}
		node.left = dfs(node.left)
		node.right = dfs(node.right)
		return res
	}
	return result
}

func main() {
	r := buildTree([]int{1,2,3,4,5,6,7,8})
	fmt.Println(printTree(r, map[int]bool{}))
}
