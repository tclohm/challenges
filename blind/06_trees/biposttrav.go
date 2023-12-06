package main

import (
	tree "trees/common"
	"fmt"
)

func postorder(n *tree.Node) []int {
	stack := []*tree.Node{n}
	result := []int{}
	visited := []bool{false}
	
	for len(stack) != 0 {
		node, v := stack[len(stack) - 1], visited[len(visited) - 1]
		stack = stack[:len(stack) - 1]
		visited = visited[:len(visited) - 1]

		if node != nil || node.Value != 0 {
			if v {
				result = append(result, node.Value)
			} else {
				stack = append(stack, node)
				visited = append(visited, true)
				
				if node.Left != nil {
					if node.Left.Value != 0 {
						stack = append(stack, node.Left)
						visited = append(visited, false)
					}
				}
				
				if node.Right != nil {
					if node.Right.Value != 0 {
						stack = append(stack, node.Right)
						visited = append(visited, false)
					}
				}
			}
		}
	}
	return result
}

func main() {
	r := tree.Build_binary_tree([]int{1,0,2,0,0,3})
	tree.Print_binary_tree(r, 0, "Root: ")
	fmt.Println(postorder(r))
}