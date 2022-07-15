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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dfs(root *TreeNode) (int, int) {
	if root != nil {
		leftDepth, maxLeft := dfs(root.Left)
		rightDepth, maxRight := dfs(root.Right)
		return max(leftDepth, rightDepth)+1, max(leftDepth+rightDepth, max(maxLeft, maxRight))
	}
	return 0,0
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, max := dfs(root)
	return max
}

func main() {
	five := &TreeNode{Val: 5}
	four := &TreeNode{Val: 4}
	two := &TreeNode{Val: 2, Left: four, Right: five}
	three := &TreeNode{Val: 3}
	one := &TreeNode{Val: 1, Left: two, Right: three}

	otherTwo := &TreeNode{Val: 2}
	otherOne := &TreeNode{Val: 1, Left: otherTwo, Right: nil}
	

	fmt.Println(__print__(one))

	fmt.Println(diameterOfBinaryTree(one))

	fmt.Println(__print__(otherOne))
	fmt.Println(diameterOfBinaryTree(otherOne))
}