package main

import (
	tree "trees/common"
	"fmt"
)

func isValid(root *tree.Node) bool {
	var dfs func(root, min, max *tree.Node) bool

	dfs = func(root, min, max *tree.Node) bool {
		if root == nil { return true }

		if min != nil && root.Value <= min.Value {
			return false
		}

		if max != nil && root.Value >= max.Value {
			return false
		}

		return dfs(root.Left, min, root) && dfs(root.Right, root, max)
	}

	return dfs(root, nil, nil)
}
func main() {
	fmt.Println(isValid(tree.Build_binary_tree([]int{1,2,3,4,5,6,7})))
	fmt.Println(isValid(tree.Build_binary_tree([]int{2,1,3})))
	fmt.Println(isValid(tree.Build_binary_tree([]int{5,1,4,0,0,3,6})))
}