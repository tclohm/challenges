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

func line_break() {
	fmt.Println("\n", strings.Repeat("-", 45), "\n")
}

func is_same_tree(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil || a.value != b.value { 
		return false 
	}

	return is_same_tree(a.left, b.left) && is_same_tree(a.right, b.right)
}

func main() {
	a := build_binary_tree([]int{1,2,3})
	b := build_binary_tree([]int{1,2,3})
	fmt.Println(is_same_tree(a, b))
	c := build_binary_tree([]int{1,2,1})
	d := build_binary_tree([]int{1,1,2})
	fmt.Println(is_same_tree(c, d))
}