package main

import (
	tree "trees/common"
	"fmt"
)

func order(root *tree.Node) [][]int {
	if root == nil { return nil }

	var result = [][]int{}
	var level = 0
	var queue = []*tree.Node{root}

	for len(queue) > 0 {
		length := len(queue)

		for i := 0 ; i < length ; i++ {
			node := queue[0]
			queue = queue[1:]

			if len(result) <= level {
				result = append(result, []int{node.Value})
			} else {
				result[level] = append(result[level], node.Value)
			}

			if node.Left != nil { queue = append(queue, node.Left) }
			if node.Right != nil { queue = append(queue, node.Right) }
		}
		level++
	}
	return result
}

func dfsOrder(root *tree.Node) [][]int {
	var result [][]int

	dfs(root, &result, 0)

	return result
}

func dfs(root *tree.Node, result *[][]int, level int) {
	if root == nil { return }

	if len(*result) == level {
		*result = append(*result, []int{})
	}

	(*result)[level] = append((*result)[level], root.Value)

	dfs(root.Left, result, level + 1)
	dfs(root.Right, result, level + 1)
}

func main() {
	r := tree.Build_binary_tree([]int{1,2,3,4,5,6,7})
	fmt.Println(order(r))
	fmt.Println(dfsOrder(r))
}