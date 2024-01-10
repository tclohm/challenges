package main

import (
	tree "trees/common"
	"fmt"
)


func similar(root1, root2 *tree.Node) bool {
	res1 := []int{}
	res2 := []int{}
	// DFS traversal
	var dfs func(node *tree.Node, leaf *[]int)
	dfs = func(node *tree.Node, leaf *[]int) {
		if nil == node { return }
		if node.Value == 0 { return }
		if nil == node.Left && nil == node.Right {
			*leaf = append(*leaf, node.Value)
			return
		}
		dfs(node.Left, leaf)
		dfs(node.Right, leaf)
	}

	dfs(root1, &res1)
	dfs(root2, &res2)

	fmt.Println(res1, res2)
	return false
}

func main() {
	t1 := tree.Build_binary_tree([]int{3,5,1,6,2,9,8,0,0,7,4,0,0})
	t2 := tree.Build_binary_tree([]int{3,5,1,6,7,4,2,0,0,0,0,0,0,9,8})
	tree.Print_binary_tree(t1, 0, "Root")
	fmt.Println("----------------------")
	tree.Print_binary_tree(t2, 0, "Root")
	fmt.Println(similar(t1, t2))
}