package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func insertLevel(arr []int, i, n int) *TreeNode {
	var root *TreeNode

	if i < n {
		root = &TreeNode{Val: arr[i]}

		root.Left = insertLevel(arr, 2 * i + 1, n)

		root.Right = insertLevel(arr, 2 * i + 2, n)
	}

	return root
}

func inOrder(root *TreeNode) {
	if root != nil {
		inOrder(root.Left)
		fmt.Print(root.Val, " ")
		inOrder(root.Right)
	}
}

func preOrder(root *TreeNode) {
	if root != nil {
		fmt.Print(root.Val, " ")
		preOrder(root.Left)
		preOrder(root.Right)
	}
}

func postOrder(root *TreeNode) {
	if root != nil {
		postOrder(root.Left)
		postOrder(root.Right)
		fmt.Print(root.Val, " ")
	}
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	length := len(values)
	var root *TreeNode
	root = insertLevel(values, 0, length)
	fmt.Println("in order")
	inOrder(root)
	fmt.Println("\npre order")
	preOrder(root)
	fmt.Println("\npost order")
	postOrder(root)
}