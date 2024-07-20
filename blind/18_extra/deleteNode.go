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

func deleteNodes(root *Node, toDelete []int) []*Node {
	del := map[int]bool{}
	for _, v := range toDelete {
		del[v] = true
	}

	forest := make([]*Node, 0, 10)

	if root != nil && !del[root.val] {
		forest = append(forest, root)
	}

	var dfs func(root *Node, d map[int]bool, array *[]*Node) *Node
	dfs = func(root *Node, d map[int]bool, array *[]*Node) *Node {
		if root == nil {
			return nil
		}
		
		if _, found := d[root.val]; found {
			if root.left != nil && !d[root.left.val] {
				*forest = append(*forest, root.left)
			}

			if root.right != nil && !d[root.right.val] {
				*forest = append(*forest, root.right)
			}

			dfs(root.left, d, forest)
			dfs(root.right, d, forest)

			return nil
		}

		root.left = dfs(root.left, d, forest)
		root.right = dfs(root.right, d, forest)
		return root
	}

	dfs(root, del, &forest)

	return forest
}

func main() {
	r := buildTree([]int{1,2,3,4,5,6,7})
	fmt.Println(printTree(r, map[int]bool{}))
	fmt.Println(deleteNodes(r, []int{3,5}))
	r2 := buildTree([]int{1,2,4,0,3})
	fmt.Println(printTree(r2, map[int]bool{}))
	fmt.Println(deleteNodes(r2, []int{3}))
}
