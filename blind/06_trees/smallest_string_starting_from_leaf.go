package main

import (
	"fmt"
	tree "trees/common"
)

// given root value [0,25] -- ['a', z']
// lexicographically smallest string at a leaf

func smallest(root *tree.Node) string {

	var result string = ""

	var dfs func(node *tree.Node, text string)
	dfs = func(node *tree.Node, text string) {
		if node == nil { return }

		text = string('a' + node.Value) + text
		if nil == node.Right && nil == node.Left {
			if result == "" || result > text {
				result = text
			}
			return
		}

		dfs(node.Left, text)
		dfs(node.Right, text)
	}

	dfs(root, "")

	return result
}


func main() {
	root := tree.Build_binary_tree([]int{0,1,2,3,4,3,4})
	tree.Print_binary_tree(root, 0, "root")
	fmt.Println(smallest(root))

	root2 := tree.Build_binary_tree([]int{25,1,3,1,3,0,2})
	tree.Print_binary_tree(root2, 0, "root")
	fmt.Println(smallest(root2))
}