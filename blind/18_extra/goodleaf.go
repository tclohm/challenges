package main

import (
	"fmt"
	tree "extra/tree"
)


func numberGood(root *tree.Node, distance int) int {
	var pairs int = 0

	var dfs func (node *tree.Node) []int
	dfs = func (node *tree.Node) []int {
		if node == nil {
			return []int{}
		}

		if nil == node.Left && nil == node.Right {
			return []int{1}
		}

		leftDistance := dfs(node.Left)
		rightDistance := dfs(node.Right)

		allDistance := append([]int{}, leftDistance...)
		allDistance = append(allDistance, rightDistance...)

		for d1 := range leftDistance {
			for d2 := range rightDistance {
				if d1 + d2 <= distance {
					pairs += 1
				}
			}
		}

		for i, _ := range allDistance {
			allDistance[i] += 1
		}

		return allDistance
	}

	dfs(root)

	return pairs
}

func main() {
	root := tree.Build_binary_tree([]int{1,2,3,0,4})
	tree.Print_binary_tree(root, 1, "")
	fmt.Println(numberGood(root, 3))
}
