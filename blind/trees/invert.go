package main

import (
	"trees/base"
)

func main() {
	r := base.Build_binary_tree([]int{1,2,3,4})
	base.Print_binary_tree(r, 0, "Root: ")
}