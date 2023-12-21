package main

import (
	tree "trees/common"
	"fmt"
)

func pathsum(root *tree.Node, target int) bool {
	if nil == root {
		return false
	}
	//fmt.Println(target, root.Value)
	// this is an awesome trick to know! start with nil first!
	if nil == root.Left && nil == root.Right {
		return root.Value == target
	}
	return pathsum(root.Left, target - root.Value) || pathsum(root.Right, target - root.Value)
}

func main() {
	t := tree.Build_binary_tree([]int{5,4,8,11,0,13,4,7,2,0,0,0,0,0,1})
	fmt.Println(pathsum(t, 22))
	fmt.Println(pathsum(t, 26))
}