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

func in(num int, array []int) bool {
	for n := range array {
		if n == num {
			return true
		}
	}
	return false
}

func deleteNodes(root *Node, toDelete []int) [][]int {
	var result = [][]int{}
	var dfs func(root *Node, path []int)
	dfs = func(root *Node, path []int) {
		if !in(root.val, toDelete) {
			path = append(path, root.val)

			if nil != root.left {
				dfs(root.left, path)
			}

			if nil != root.right {
				dfs(root.right, path)
			}

			result = append(result, path)
		}
	}

	dfs(root, []int{})
	return result
}

func main() {
	r := buildTree([]int{1,2,3,4,5,6,7})
	fmt.Println(printTree(r, map[int]bool{}))
	fmt.Println(deleteNodes(r, []int{3,5}))
	r2 := buildTree([]int{1,2,4,0,3})
	fmt.Println(printTree(r2, map[int]bool{}))
	fmt.Println(deleteNodes(r2, []int{3}))
}
