package main

import (
	"fmt"
	"math"
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

func getDepth(root *TreeNode) int {
	if root != nil {
		left := getDepth(root.Left)
		right := getDepth(root.Right)

		if left > right {
			return left + 1
		}
		return right + 1
	}
	return 0
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}

	left, right := getDepth(root.Left), getDepth(root.Right)

	if math.Abs(float64(left - right)) > 1 {
		return false
	}

	return true
}

func main() {
	nine := &TreeNode{Val: 9}
	fifteen := &TreeNode{Val: 15}
	seven := &TreeNode{Val: 7}
	twenty := &TreeNode{Val: 3, Left: fifteen, Right: seven}
	three := &TreeNode{Val: 3, Left: nine, Right: twenty}

	fmt.Println(__print__(three))

	fmt.Println(isBalanced(three))

	fourR := &TreeNode{Val: 4}
	fourL := &TreeNode{Val: 4}
	threeL := &TreeNode{Val: 3, Left: fourL, Right: fourR}
	threeR := &TreeNode{Val: 3}
	twoL := &TreeNode{Val: 2, Left: threeL, Right: threeR}
	twoR := &TreeNode{Val: 2}
	rootOne := &TreeNode{Val: 1, Left: twoL, Right: twoR}

	fmt.Println(__print__(rootOne))
	fmt.Println(isBalanced(rootOne))
}