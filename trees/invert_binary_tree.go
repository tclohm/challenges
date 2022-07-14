package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) {
	if root == nil { return }

	if root.Left  != nil { dfs(root.Left)  }
	if root.Right != nil { dfs(root.Right) }

	root.Left, root.Right = root.Right, root.Left
}

func invert(root *TreeNode) *TreeNode {
	// recursive, using comps built in stack
	// time: O(N)
	dfs(root)
	return root
}

// iterative
func invertAgain(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		leaf := stack[len(stack) - 1]
		stack = stack[0:len(stack) - 1]

		if leaf.Left != nil {
			stack = append(stack, leaf.Left)
		}

		if leaf.Right != nil {
			stack = append(stack, leaf.Right)
		}

		leaf.Left, leaf.Right = leaf.Right, leaf.Left
	}

	return root
}

func __print__(root *TreeNode) []int {
	represent := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		r := queue[0]
		represent = append(represent, r.Val)

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

func main() {
	r := &TreeNode{Val: 3}
	l := &TreeNode{Val: 1}
	root := &TreeNode{Val: 2, Right: r, Left: l}
	fmt.Println(__print__(root))
	invert(root)
	fmt.Println(__print__(root))

	fmt.Println("---------")

	nine := &TreeNode{Val: 9}
	six := &TreeNode{Val: 6}
	seven := &TreeNode{Val: 7, Left: six, Right: nine}

	three := &TreeNode{Val: 3}
	one := &TreeNode{Val: 1}
	two := &TreeNode{Val: 2, Left: one, Right: three}

	four := &TreeNode{Val: 4, Left: two, Right: seven}

	fmt.Println(__print__(four))
	invert(four)
	fmt.Println("---------")
	fmt.Println("Flip!")
	fmt.Println(__print__(four))
	fmt.Println("---------")
	fmt.Println("Flip back")
	invertAgain(four)
	fmt.Println(__print__(four))
}