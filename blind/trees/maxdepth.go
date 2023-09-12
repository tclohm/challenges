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

func maxdepth(root *Node) int {
	deepest := 0
	var dfs func(root *Node, level int, deepest *int)
	dfs = func(root *Node, level int, deepest *int)  {
		if level > *deepest {
			*deepest = level
		}
		if root.left != nil {
			dfs(root.left, level+1, deepest)
		}
		if root.right != nil {
			dfs(root.right, level+1, deepest)
		}
	}

	dfs(root, 0, &deepest)
	return deepest + 1
}

func line_break() {
	fmt.Println("\n", strings.Repeat("-", 45), "\n")
}


func main() {
	three := build_binary_tree([]int{3,9,20,-1,-1,15,7})
	print_binary_tree(three, 0, "Root: ")
	fmt.Println(maxdepth(three))
	line_break()
	two := build_binary_tree([]int{1,-1,2})
	print_binary_tree(two, 0, "Root: ")
	fmt.Println(maxdepth(two))

}