package main

import (
	"fmt"
	tree "trees/common"
)

func binaryTreePaths(root *tree.Node) []string {
	result := []string{}

	var dfs func (node *tree.Node, res []string)
	dfs = func(node *tree.Node, res []string) {
		res = append(res, fmt.Sprintf("%d", node.Value))
		if node.Left == nil && node.Right == nil {
			result = append(result, res...)
		}

		res = append(res, "->")
		if node.Left != nil {
			dfs(node.Left, res)
		}
		if node.Right != nil {
			dfs(node.Right, res)
		}
	}

	dfs(root, nil)
	return result
}

func main() {
	root := tree.Build_binary_tree([]int{1,2,3,0,5})
	tree.Line_break()
	fmt.Println(binaryTreePaths(root))
}