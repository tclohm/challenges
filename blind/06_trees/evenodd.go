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

func isEvenOddTree(result [][]int) bool {
	for i, array := range result {
		if i % 2 == 0 {
			for _, value := range array {
				if value % 2 == 0 {
					return false
				}
			}
		} else {
			for _, value := range array {
				if value % 2 == 1 {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	ex1 := tree.Build_binary_tree([]int{1,10,4,3,0,7,9,12,8,6,0,0,2})
	ex2 := tree.Build_binary_tree([]int{5,4,2,3,3,7})
	ex3 := tree.Build_binary_tree([]int{5,9,1,3,5,7})
	fmt.Println(isEvenOddTree(order(ex1)))
	fmt.Println(isEvenOddTree(order(ex2)))
	fmt.Println(isEvenOddTree(order(ex3)))
}