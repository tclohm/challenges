package main

import "fmt"

type Node struct {
	val int
	left *Node
	right *Node
}

func numberGood(root *Node, distance int) int {
	return 0
}

func main() {
	root := []int{1,2,3,0,4}
	fmt.Println(numberGood(root, 3))
}
