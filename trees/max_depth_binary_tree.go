package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func __print__(root *TreeNode) []string {
	represent := []string{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		r := queue[0]
		represent = append(represent, strconv.Itoa(r.Val))

		if r.Left != nil {
			queue = append(queue, r.Left)
		}

		if r.Right != nil {
			queue = append(queue, r.Right)
		}

		queue = queue[1:len(queue)]
	}

	return represent
}

// func createTree(array []int) *TreeNode {
// 	// left index + 1
// 	// right index + 2

// 	var head *TreeNode = &TreeNode{Val: array[0]}

// 	for index := 0 ; index < len(array) ; index++ {
// 		var node *TreeNode = &TreeNode{node}
// 		var left, right *TreeNode

// 		if index + 1 <= len(array) - 1 {

// 		}

// 		if index + 2 <= len(array) - 1 {
			
// 		}
// 	}
// }

type Data struct {
	Node *TreeNode
	Level int
}

func bfs(root *TreeNode, level int) int {
	q := make([]Data, 0)
	d := Data{Node: root, Level: level}
	q = append(q, d)

	depth := 0

	for len(q) > 0 {
		data := q[0]
		q = q[1:len(q)]

		if data.Level > depth {
			depth = data.Level
		}

		if data.Node.Left != nil {
			d := Data{Node: data.Node.Left, Level: data.Level + 1}
			q = append(q, d)
		}

		if data.Node.Right != nil {
			d := Data{Node: data.Node.Right, Level: data.Level + 1}
			q = append(q, d)
		}
	}

	return depth
}

func findDepth(root *TreeNode) int {
	height := bfs(root, 0)
	return height + 1
}

func main() {
	seven := &TreeNode{Val: 7, Left: nil, Right: nil}
	fifteen := &TreeNode{Val: 15, Left: nil, Right: nil}
	twenty := &TreeNode{Val: 20, Left: fifteen, Right: seven}
	nine := &TreeNode{Val: 9, Left: nil, Right: nil}
	root := &TreeNode{Val: 3, Left: nine, Right: twenty}

	fmt.Println(__print__(root))

	fmt.Println(findDepth(root))
}