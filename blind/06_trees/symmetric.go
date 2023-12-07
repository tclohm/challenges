package main

import (
	tree "trees/common"
	"fmt"
)

func symmetric(n *tree.Node) bool {
	var dfs func (left, right *tree.Node) bool
	dfs = func (left, right *tree.Node) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}

		return (left.Value == right.Value && dfs(left.Left, right.Right) && dfs(left.Right, right.Left))
	}

	return dfs(n.Left, n.Right)
}

func main() {
	one := tree.Build_binary_tree([]int{1,2,2,3,4,4,3})
	tree.Print_binary_tree(one, 0, "Root")
	fmt.Println(symmetric(one))
	fmt.Println("-------------")
	two := tree.Build_binary_tree([]int{1,2,3,4,5,6})
	tree.Print_binary_tree(two, 0, "Root")
	fmt.Println(symmetric(two))
}