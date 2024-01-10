package main

import (
	tree "trees/common"
	"fmt"
)

func main() {
	t1 := tree.Build_binary_tree([]int{3,5,1,6,2,9,8,0,0,7,4,0,0})
	t2 := tree.Build_binary_tree([]int{3,5,1,6,7,4,2,0,0,0,0,0,0,9,8})
	tree.Print_binary_tree(t1, 0, "Root")
	fmt.Println("----------------------")
	tree.Print_binary_tree(t2, 0, "Root")
}