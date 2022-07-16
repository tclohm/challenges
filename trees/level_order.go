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

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{{root.Val}}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		r := queue[0]

		level := []int{}

		if r.Left != nil {
			queue = append(queue, r.Left)
			level = append(level, r.Left.Val)
		}

		if r.Right != nil {
			queue = append(queue, r.Right)
			level = append(level, r.Right.Val)
		}

		if len(level) > 0 {
			result = append(result, level)
		}

		queue = queue[1:len(queue)]
	}

	return result
}

func main() {
	fifteen := &TreeNode{Val: 15}
	seven := &TreeNode{Val: 7}
	twenty := &TreeNode{Val: 20, Left: fifteen, Right: seven}
	nine := &TreeNode{Val: 9}
	three := &TreeNode{Val: 3, Left: nine, Right: twenty}

	fmt.Println(levelOrder(three))

}