package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, answer *int, maxValueSeen int) {
	if root == nil {
		return
	}

	if root.Val >= maxValueSeen {
		*answer++
		maxValueSeen = root.Val
	}

	dfs(root.Left, answer, maxValueSeen)
	dfs(root.Right, answer, maxValueSeen)
}

func goodNodes(root *TreeNode) int {
	// node X in the tree is named good if the path from root to X there is no greater value than x
	answer := 0
	dfs(root, &answer, root.Val)
	return answer
}

func main() {
	five := &TreeNode{Val: 5}
	one := &TreeNode{Val: 1}
	otherThree := &TreeNode{Val: 3}
	otherOne := &TreeNode{Val: 1, Left: otherThree}
	four := &TreeNode{Val: 4, Left: one, Right: five}
	three := &TreeNode{Val: 3, Left: otherOne, Right: four}
	

	fmt.Println(goodNodes(three))
}