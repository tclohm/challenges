package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func __print__(root *TreeNode) []string {
	represent := []string{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		r := queue[0]
		represent = append(represent, strconv.Itoa(r.Val))

		if r.Left != nil {
			queue = append(queue, r.Left)
		}

		if r.Right != nil {
			queue = append(queue, r.Right)
		}

		queue = queue[1:len(queue)]
	}

	return represent
}

func isSameTree(root, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == nil && subRoot == nil
	}

	return root.Val == subRoot.Val && isSameTree(root.Left, subRoot.Left) && isSameTree(root.Right, subRoot.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil { return subRoot == nil }

	return isSameTree(root, subRoot) || isSameTree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func main() {
	two := &TreeNode{Val: 2}
	one := &TreeNode{Val: 1}
	subRoot := &TreeNode{Val: 4, Left: one, Right: two}

	zero := &TreeNode{Val: 0}
	two1 := &TreeNode{Val: 2, Left: zero}
	one1 := &TreeNode{Val: 1}
	four := &TreeNode{Val: 4, Left: one1, Right: two1}
	five := &TreeNode{Val: 5}
	root := &TreeNode{Val: 3, Left: four, Right: five}

	fmt.Println(isSubtree(root, subRoot))
}