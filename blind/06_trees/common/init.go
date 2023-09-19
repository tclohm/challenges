package common

import (
	"fmt"
	"strings"
)

type Node struct {
	Value int
	Left *Node
	Right *Node
}

func Build_binary_tree(list []int) *Node {
	if len(list) == 0 { return nil }
	nodes := make([]*Node, len(list))

	for i := 0 ; i < len(list) ; i++ {
		nodes[i] = &Node{Value: list[i]}
	}

	root := nodes[0]
	index := 0

	for _, node := range nodes {
		if node != nil {
			left_index := 2 * index + 1
			right_index := 2 * index + 2

			if left_index < len(list) {
				node.Left = nodes[left_index]
			}

			if right_index < len(list) {
				node.Right = nodes[right_index]
			}

			index += 1
		}
	}

	return root
}

func Print_binary_tree(root *Node, level int, prefix string) {
	if root != nil {
		fmt.Println(strings.Repeat(" ", level * 4), prefix, root.Value)
		if root.Left != nil || root.Right != nil {
			Print_binary_tree(root.Left, level+1, "L--- ")
			Print_binary_tree(root.Right, level+1, "R--- ")
		}
	}
}

func Line_break() {
	fmt.Println("\n", strings.Repeat("-", 45), "\n")
}