package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, isValid *bool) {
	if root == nil {
		return
	}

	if root.Right != nil && root.Right.Val < root.Val || root.Left != nil && root.Left.Val > root.Val {
		*isValid = false
	}

	dfs(root.Left, isValid)
	dfs(root.Right, isValid)

}

func isValidBST(root *TreeNode) bool {
	var isValid bool = true
	dfs(root, &isValid)
	return isValid
}


func main() {
	six := &TreeNode{Val: 6}
	three := &TreeNode{Val: 3}

	four := &TreeNode{Val: 4, Left: three, Right: six}
	one := &TreeNode{Val: 1}

	five := &TreeNode{Val: 5, Left: one, Right: four}

	fmt.Println(isValidBST(five))

}