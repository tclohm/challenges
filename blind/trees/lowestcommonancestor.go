package main

import (
	tree "trees/common"
	"fmt"
)

// root is the most common ancestor, maybe not the lowest
func lca(root, n1, n2 *tree.Node) *tree.Node {
	// right is bigger and left is lower
	if n1.Value > root.Value && n2.Value > root.Value {
		return lca(root.Right, n1, n2)
	}

	if n1.Value < root.Value && n2.Value < root.Value {
		return lca(root.Left, n1, n2)
	}

	return root
}

func main() {
	r := tree.Build_binary_tree([]int{1,2,3,4,5,6,7})
	// ugh I'm not checking this
}