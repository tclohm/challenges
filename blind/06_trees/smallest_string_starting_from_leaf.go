package main

import (
	"fmt"
	tree "trees/common"
)

// given root value [0,25] -- ['a', z']
// lexicographically smallest string at a leaf

func smallest(root *tree.Node) {
	stack := []*tree.Node{root}

	for len(stack) > 0 {
		end := len(stack) - 1
		node := stack[end]
		stack = stack[:end]

		fmt.Println(string('a' + node.Value))

		if nil != node.Left {
			stack = append(stack, node.Left)
		}

		if nil != node.Right {
			stack = append(stack, node.Right)
		}
	}
}


func main() {
	root := tree.Build_binary_tree([]int{0,1,2,3,4,3,4})
	tree.Print_binary_tree(root, 0, "root")
	smallest(root)
}