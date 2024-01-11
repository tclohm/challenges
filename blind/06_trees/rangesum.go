package main

import (
	tree "trees/common"
	"fmt"
)

func inorder(root *tree.Node) []int {
	var array = []int{}
	var dfs func(root *tree.Node)
	dfs = func(root *tree.Node) {
		if nil == root { return }
		if root.Value == 0 { return }
		if nil != root.Left {
			dfs(root.Left)
		}
		array = append(array, root.Value)
		if nil != root.Right {
			dfs(root.Right)
		}
	}
	dfs(root)
	return array
}

// func rangeSum(root *tree.Node, low, high int) int {
// 	array := inorder(root)
// 	total := 0 
// 	for i := 0 ; i < len(array) ; i++ {
// 		if array[i] >= low && array[i] <= high {
// 			total += array[i]
// 		}
// 	}
// 	return total
// }

func rangeSum(root *tree.Node, low, high int) int {
	if nil == root || root.Value == 0 {
		return 0
	}

	if root.Value > high {
		return rangeSum(root.Left, low, high)
	}
	if root.Value < low {
		return rangeSum(root.Right, low, high) 
	}
	return root.Value + rangeSum(root.Left, low, high) + rangeSum(root.Right, low, high)
}


func main() {
	t1 := tree.Build_binary_tree([]int{10,5,15,3,7,0,18})
	tree.Print_binary_tree(t1, 0, "Root")
	// fmt.Println("----------------------")
	// tree.Print_binary_tree(t2, 0, "Root")
	fmt.Println(rangeSum(t1, 7, 15))
}