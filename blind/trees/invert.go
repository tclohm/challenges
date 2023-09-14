package main

import (
	tree "trees/common"
	"fmt"
)

func invert(root *tree.Node) *tree.Node {
	if root == nil { return root }

	left := invert(root.Right)
	root.Right = invert(root.Left)
	root.Left = left

	return root

}

func main() {
	r := tree.Build_binary_tree([]int{1,2,3,4,5,6,7})
	tree.Print_binary_tree(r, 0, "Root: ")
	tree.Line_break()
	fmt.Println("inverted")
	tree.Print_binary_tree(invert(r), 0, "Root :")
}