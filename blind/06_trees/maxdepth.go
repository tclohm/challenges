package main

import (
	tree "trees/common"
	"fmt"
)

func maxdepth(root *tree.Node) int {
	// right + 1, left + 1
	if root == nil { return 0 }

	left := maxdepth(root.Left) 
	right := maxdepth(root.Right)
	
	if left > right { return left + 1 }
	return right + 1
}

func main() {
	r := tree.Build_binary_tree([]int{1,2,3,4,5,6,7})
	tree.Print_binary_tree(r, 0, "Root: ")
	fmt.Println("max depth:", maxdepth(r))
	tree.Line_break()
	l := tree.Build_binary_tree([]int{1,2,3,4,5,6,7,8,9,10,12,13,14,15,16,17})
	tree.Print_binary_tree(l, 0, "Root: ")
	fmt.Println("max depth:", maxdepth(l))

}