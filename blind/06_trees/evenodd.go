package main

import (
	tree "trees/common"
	"fmt"
)


func isEvenOddTree(root *tree.Node) bool {
	if root == nil { return true }

	var even = true
	var queue = []*tree.Node{root}

	for len(queue) > 0 {
		var prev int
		if even {
			prev = -10000
		} else {
			prev = 10000
		}
		length := len(queue)

		for i := 0 ; i < length ; i++ {
			node := queue[0]
			queue = queue[1:]

			if even && (node.Value % 2 == 0 || node.Value <= prev) {
				return false
			} else if !even && (node.Value % 2 == 1 || node.Value >= prev) {
				return false
			}

			if node.Left != nil { queue = append(queue, node.Left) }
			if node.Right != nil { queue = append(queue, node.Right) }
			prev = node.Value
		}
		even = !even
	}
	return true
}

func main() {
	ex1 := tree.Build_binary_tree([]int{1,10,4,3,0,7,9,12,8,6,0,0,2})
	ex2 := tree.Build_binary_tree([]int{5,4,2,3,3,7})
	ex3 := tree.Build_binary_tree([]int{5,9,1,3,5,7})
	fmt.Println(isEvenOddTree(ex1))
	fmt.Println(isEvenOddTree(ex2))
	fmt.Println(isEvenOddTree(ex3))
}