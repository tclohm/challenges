package main

import (
	tree "trees/common"
	"fmt"
)


func preorder(n *tree.Node) []int {
	result := []int{}
	queue := []*tree.Node{n}

	for len(queue) != 0 {
		
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Value)

		if node.Left != nil {
			if node.Left.Value != 0 {
				queue = append(queue, node.Left)
			}
		}
		if node.Right != nil {
			if node.Right.Value != 0 {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}

func main() {
	r := tree.Build_binary_tree([]int{1,0,2,0,0,3})
	tree.Print_binary_tree(r, 0, "Root: ")
	fmt.Println(preorder(r))
}