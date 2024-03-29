package main

import (
	"fmt"
	"math"
)

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	depth++
	return max(maxDepth(root.Left, depth), maxDepth(root.Right, depth))
}

type doublearray [][]int

func depths(root *TreeNode, depth int, levels *doublearray) {
	if root == nil { return }
	//fmt.Println("levels", depth, levels)
	if len(*levels) > depth {
		(*levels)[depth] = append((*levels)[depth], root.Val)
	} else {
		(*levels) = append((*levels), []int{root.Val})
	}
	depth++
	depths(root.Left, depth, levels)
	depths(root.Right, depth, levels)
}

func levels(root *TreeNode) [][]int {
	var result doublearray = [][]int{}
	depths(root, 0, &result)
	return result
}

type Queue struct {
	Container []*TreeNode
}

func (q *Queue) Push(item *TreeNode) {
	q.Container = append(q.Container, item)
}

func (q *Queue) Pop() *TreeNode {
	if len(q.Container) == 0 { return nil }
	node := q.Container[0]
	q.Container = q.Container[1:]
	return node
}

func (q *Queue) Length() int {
	return len(q.Container)
}

type Stack struct {
	Container []*TreeNode
}

func (s *Stack) Push(item *TreeNode) {
	s.Container = append(s.Container, item)
}

func (s *Stack) Pop() *TreeNode {
	if len(s.Container) == 0 { return nil }
	length := len(s.Container) - 1
	node := s.Container[length]
	s.Container = s.Container[:length]
	return node
}

func (s *Stack) Length() int {
	return len(s.Container)
}

func bfs(root *TreeNode) doublearray {
	if root == nil { return [][]int{} }
	
	var result doublearray = [][]int{}
	var q = Queue{}

	q.Push(root)

	for q.Length() != 0 {

		var length, count = q.Length(), 0
		var currentLevel = []int{}

		for count < length {

			var node = q.Pop()
			currentLevel = append(currentLevel, node.Val)
			if node.Left != nil { q.Push(node.Left) }
			if node.Right != nil { q.Push(node.Right) }
			count++
			
		}

		result = append(result, currentLevel)
	}
	return result
}

						// 		     1
						// 	    2          3
						// 	 4    5     6     7
						// 	8 9 10 11 12 13 14 15
						// 16

func rstBfs(root *TreeNode) []int {
	if root == nil { return []int{} }
	var result = []int{}
	var q = Queue{}
	q.Push(root)

	for q.Length() != 0 {

		var length, count = q.Length(), 0 

		for count < length {
			var node = q.Pop()
			if node.Left != nil { q.Push(node.Left) }
			if node.Right != nil { q.Push(node.Right) }
			count++
			if count == length { result = append(result, node.Val) }
		}
	}
	return result
}

func dfsRightSide(root *TreeNode) []int {
	var result = []int{}
	dfs(root, 0, &result)
	return result
}

func dfs(root *TreeNode, currentLevel int, result *[]int) {
	// go right first as far as we can go
	// save the depth and values
	// then save depth left if there is no right
	if root == nil { return }
	if currentLevel >= len(*result) {
		*result = append(*result, root.Val)
	}
	if root.Right != nil {
		dfs(root.Right, currentLevel + 1, result)
	}
	if root.Left != nil {
		dfs(root.Left, currentLevel + 1, result)
	}
}

func countNodes(root *TreeNode) int {
	if root == nil { return 0 }
	var height = getTreeHeight(root)

	if height == 0 { return 1 }

	var upperbound = math.Floor(math.Pow(2, float64(height)))
	var left, right = 0, upperbound

	for left < int(right) {
		var indexToFind = left + (int(right) - left) / 2
		if nodeExists(indexToFind, height, root) { 
			left =  indexToFind 
		} else {
			right = float64(indexToFind - 1)
		}
	}
	return int(right)
}

func getTreeHeight(root *TreeNode) int {
	var height = 0
	for root.Left != nil {
		height += 1
		root = root.Left
	}
	return height
}

func nodeExists(indexToFind, height int, node *TreeNode) bool {
	var left, count = 0, 0
	var right = math.Floor(math.Pow(2, float64(height)))
	for count < height {
		var midOfNode = left + (int(right) - left) / 2
		if indexToFind >= midOfNode {
			node = node.Right
			left = midOfNode
		} else {
			node = node.Left
			right = float64(midOfNode - 1)
		}
		count += 1
	}

	return node != nil
}

func dfsValid(root *TreeNode, min, max int) bool {
	if root.Val <= min || root.Val >= max { return false }

	if root.Left != nil {
		if !dfsValid(root.Left, min, root.Val) { return false }
	}

	if root.Right != nil {
		if !dfsValid(root.Right, root.Val, max) { return false }
	}

	return true
}

func isValidBST(root *TreeNode) bool {
	if root == nil { return true }
	return dfsValid(root, -1000000, 1000000)
}


func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	length := len(values)
	var root *TreeNode
	root = insertLevel(values, 0, length)
	fmt.Println("in order")
	inOrder(root)
	fmt.Println("\npre order")
	preOrder(root)
	fmt.Println("\npost order")
	postOrder(root)
	fmt.Println("\nmax depth:", maxDepth(root, 0))

	fmt.Println("\nlevels recursion", levels(root))

	fmt.Println("\nlevels bfs", bfs(root))

	fmt.Println("right side bfs", rstBfs(root))

	fmt.Println("right side dfs", dfsRightSide(root))

	fmt.Println("-------")

	v := []int{12, 7, 18, 5, 9, 16, 25}
	r := insertLevel(v, 0, len(v))

	fmt.Println(isValidBST(r))
}