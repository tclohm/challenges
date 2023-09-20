package main

import (
	tree "trees/common"
)

// first value in preorder will always have the root
// inorder will give us the subarrays of the root
func buildTree(preorder []int, inorder []int) *tree.Node {
	// base case
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &tree.Node{Value: preorder[0]}
	// find position of node in order
	mid := index(inorder, preorder[0])
	// recursive left, subarrays (preorder 1, mid+1), (inorder :mid)
	root.Left = buildTree(preorder[1: mid + 1], inorder[:mid])
	// recursive right, mid + 1 for preorder, inorder mid + 1
	root.Right = buildTree(preorder[mid + 1:], inorder[mid + 1:])
	return root
}

func index(arr []int, val int) int {
	for i, v := range arr {
		if v == val { return i }
	}
	return -1
}

func main() {
	tree.Print_binary_tree(buildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7}), 0, "Root: ")
	tree.Print_binary_tree(buildTree([]int{-1}, []int{-1}), 0, "Root: ")
}