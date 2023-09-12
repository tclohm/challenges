package main

import (
	"fmt"
	"strings"
)

type Node struct {
	value int
	left *Node
	right *Node
}

func build_binary_tree(list []int) *Node {
	if len(list) == 0 { return nil }
	nodes := make([]*Node, len(list))

	for i := 0 ; i < len(list) ; i++ {
		nodes[i] = &Node{value: list[i]}
	}

	root := nodes[0]
	index := 0

	for _, node := range nodes {
		if node != nil {
			left_index := 2 * index + 1
			right_index := 2 * index + 2

			if left_index < len(list) {
				node.left = nodes[left_index]
			}

			if right_index < len(list) {
				node.right = nodes[right_index]
			}

			index += 1
		}
	}

	return root
}

func print_binary_tree(root *Node, level int, prefix string) {
	if root != nil {
		fmt.Println(strings.Repeat(" ", level * 4), prefix, root.value)
		if root.left != nil || root.right != nil {
			print_binary_tree(root.left, level+1, "L--- ")
			print_binary_tree(root.right, level+1, "R--- ")
		}
	}
}

func invert(root *Node) *Node {
	if root == nil { return root }
	
	left := invert(root.left)

	root.left = invert(root.right)
	root.right = left

	return root
}

func line_break() {
	fmt.Println("\n", strings.Repeat("-", 45), "\n")
}

func main() {
	// n
	// left = parent_index * 2 + 1
	// right = parent_index * 2 + 2
	root := build_binary_tree([]int{4, 2, 7, 1, 3, 6, 9})
	print_binary_tree(root, 0, "Root: ")
	line_break()
	invert(root)
	fmt.Println("INVERT")
	print_binary_tree(root, 0, "Root: ")
	line_break()
	r := build_binary_tree([]int{2, 1, 3})
	print_binary_tree(r, 0, "Root: ")
	invert(r)
	line_break()
	fmt.Println("INVERT")
	print_binary_tree(r, 0, "Root: ")


}