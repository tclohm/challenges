package main

import (
	tree "trees/common"
	"fmt"
)

func sametree(a, b *tree.Node) bool {
	// is a and b nil, equals true
	// if a or b is nil, or the values are not the same, equal false
	if a == nil && b == nil { return true }
	if a == nil || b == nil || a.Value != b.Value { return false }
	return sametree(a.Left, b.Left) && sametree(a.Right, b.Right)
}

func main() {
	one := tree.Build_binary_tree([]int{1,2,3,4,5,6,7})
	two := tree.Build_binary_tree([]int{1,2,3,4,5,6,7})
	three := tree.Build_binary_tree([]int{1,2,3,4,5,6,7,8})
	fmt.Println(sametree(one, two))
	fmt.Println(sametree(two, three))
}