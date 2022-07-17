package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		k -= 1
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
	return -1
}

func main() {
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2, Left: one}
	four := &TreeNode{Val: 4}
	three := &TreeNode{Val: 3, Left: two, Right: four}
	six := &TreeNode{Val: 6}
	root := &TreeNode{Val: 5, Left: three, Right: six}

	fmt.Println(kthSmallest(root, 6))


}