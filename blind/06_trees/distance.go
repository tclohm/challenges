package main

import (
	tree "trees/common"
	"fmt"
)

func distance(root tree.Node) int {
	var prev tree.Node
	var res int = 100000000

	var dfs func (node tree.Node)
	dfs = func (node tree.Node) {
		if node.Value == 0 {
			return
		}

		if node.Left != nil {
			dfs(*node.Left)
		}

		if prev.Value > 0 {
			res = min(res, node.Value - prev.Value)
		}

		prev = node

		if node.Right != nil {
			dfs(*node.Right)
		}
		
	}

	dfs(root)
	return res
}

func min(a, b int) int {
	if a < b { return a }
	return b
}

func main() {
	r := tree.Build_binary_tree([]int{4,2,6,1,3,0,0})
	tree.Print_binary_tree(r, 0, "Root:")
	fmt.Println(distance(*r))
}