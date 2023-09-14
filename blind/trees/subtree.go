package main

import (
	tree "trees/common"
	"fmt"
)

func subtree(a, b *tree.Node) bool {
	if a == nil { return false }
	
	return equals(a, b) || subtree(a.Left, b) && subtree(a.Right, b)
}

func equals(a, b *tree.Node) bool {
	if a == nil && b == nil { return true }
	if a == nil || b == nil || a.Value != b.Value { return false }
	return equals(a.Left, b.Left) && equals(a.Right, b.Right)
}

func main() {
	one := tree.Build_binary_tree([]int{1,2,3,4,5,6,7})
	two := tree.Build_binary_tree([]int{1,2,3,4,5,6,7})
	three := tree.Build_binary_tree([]int{1,2,3,4,5,6,7,8})
	fmt.Println(subtree(one, two))
	fmt.Println(subtree(two, three))
}