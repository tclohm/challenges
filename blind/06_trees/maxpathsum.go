package main

import (
	tree "trees/common"
	"fmt"
)
// find path with the tree where the sum of nodes values is max
// follow the parent-child relationship
// define a recursive function, dfs
// return two values, max path sum that can be achieved starting from the current node
// and the max path sum that can end at the current node
func sum(root *tree.Node) int {
	if root == nil { return 0 }

	max_path_sum := -10

	// recursive function to calculate max path sum starting and ending at the current node
	var dfs func(root *tree.Node) int 
	dfs = func(root *tree.Node) int {
		if root == nil { return 0 }
		// calculate max path sums for the left and right subtrees
		left_sum := max(dfs(root.Left), 0)
		right_sum := max(dfs(root.Right), 0)
		// calculate the max path sum ending at the current node
		max_end := root.Value + max(left_sum, right_sum)
		// update the global maximum path sum at the current node
		max_path_sum = max(max_path_sum, root.Value + left_sum + right_sum)
		// return the max path sum ending at the current node
		return max_end
	}

	dfs(root)

	return max_path_sum

}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func main() {
	fmt.Println(sum(tree.Build_binary_tree([]int{1,2,3})))
	fmt.Println(sum(tree.Build_binary_tree([]int{-1,2,3,4,5})))
}