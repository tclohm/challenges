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
		length := len(queue)
		var prev *tree.Node = nil

		for i := 0 ; i < length ; i++ {
			node := queue[i]

			if even {
				if prev != nil && node.Value <= prev.Value && prev.Value != 0 && node.Value != 0 {
					return false
				}
				if node.Value % 2 == 0 && node.Value != 0  {
					return false
				}
			} else {
				if prev != nil && node.Value >= prev.Value && prev.Value != 0 && node.Value != 0  {
					return false
				}

				if node.Value % 2 != 0 && node.Value != 0  {
					return false
				}
				
			}

			prev = node

			if node.Left != nil { queue = append(queue, node.Left) }
			if node.Right != nil { queue = append(queue, node.Right) }
			
		}
		queue = queue[length:]
		even = !even
	}
	return true
}

func main() {
	ex1 := tree.Build_binary_tree([]int{1,10,4,3,0,7,9,12,8,0,0,6,0,0,2})
	tree.Print_binary_tree(ex1, 0, "Root")
	ex2 := tree.Build_binary_tree([]int{5,4,2,3,3,7})
	ex3 := tree.Build_binary_tree([]int{5,9,1,3,5,7})
	fmt.Println(isEvenOddTree(ex1))
	fmt.Println(isEvenOddTree(ex2))
	fmt.Println(isEvenOddTree(ex3))
}