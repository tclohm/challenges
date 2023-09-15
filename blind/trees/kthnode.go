package main

import (
	tree "trees/common"
	"fmt"
)

func kthSmallest(root *tree.Node, k int) int {
	stack := make([]*tree.Node, 0, k)

	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		k--
		if k == 0 { return root.Value }

		root = root.Right
	}

	return -1
}

func main() {
	one := tree.Build_binary_tree([]int{3,1,4,0,2})
	six := tree.Build_binary_tree([]int{5,3,6,2,4,0,0,1})
	fmt.Println(kthSmallest(one, 1))
	fmt.Println(kthSmallest(six, 3))
}