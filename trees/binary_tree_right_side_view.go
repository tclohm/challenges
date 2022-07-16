package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		length := len(queue)

		for length > 0 {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if length == 1 {
				// the most right node on that level
				result = append(result, node.Val)
			}

			length -= 1
		}
	}

	return result
}

func main() {
	four := &TreeNode{Val: 4}
	three := &TreeNode{Val: 3, Right: four}
	five := &TreeNode{Val: 5}
	two := &TreeNode{Val: 2, Right: five}
	one := &TreeNode{Val: 1, Left: two, Right: three}

	fmt.Println(rightSideView(one))
}