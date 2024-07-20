package main

import "fmt"

type Node struct {
	val int
	left *Node
	right *Node
}

func numberGood(root *Node, distance int) int {
	var pairs int = 0

	var dfs func (node *Node) []int
	dfs = func (node *Node) []int {
		if node == nil {
			return []int{}
		}

		if nil == node.left && nil == node.right {
			return []int{1}
		}

		leftDistance := dfs(node.left)
		rightDistance := dfs(node.right)

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
	root := []int{1,2,3,0,4}
	fmt.Println(numberGood(root, 3))
}
