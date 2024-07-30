package main

import "fmt"

type Node struct {
	value int
	left *Node
	right *Node
}

func diameter(root *Node) int {
	if root == nil { return 0 }
	var left, right int
	if nil != root.left {
		left += diameter(root.left)
	}
	if nil != root.right {
		right += diameter(root.right)
	}
	return 1 + max(right, left)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	five := &Node{value: 5}
	four := &Node{value: 4}
	two := &Node{value: 2}
	two.left = four
	two.right = five
	three := &Node{value: 3}
	root := Node{value: 1}
	root.right = three
	root.left = two

	fmt.Println(diameter(&root))
}
