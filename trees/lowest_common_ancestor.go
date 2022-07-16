package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func lca(root, p, q *TreeNode) *TreeNode {
	queue := []*TreeNode{}

	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]

		if p.Val < node.Val && node.Val > q.Val {
			queue = append(queue, node.Left)
		}

		if p.Val > node.Val && node.Val < q.Val {
			queue = append(queue, root.Right)
		}

		if p.Val <= node.Val && node.Val <= q.Val {
			return node
		}

		queue = queue[1:len(queue)]
	}
	return nil
}

func main() {
	three := &TreeNode{Val: 3}
	five := &TreeNode{Val: 5}
	four := &TreeNode{Val: 4, Left: three, Right: five}
	zero := &TreeNode{Val: 0}
	two := &TreeNode{Val: 2, Left: zero, Right: four}

	seven := &TreeNode{Val: 7}
	nine := &TreeNode{Val: 9}
	eight := &TreeNode{Val: 8, Left: seven, Right: nine}
	six := &TreeNode{Val: 6, Left: two, Right: eight}

	node := lowestCommonAncestor(six, two, eight)
	n := lca(six, two, eight)
	fmt.Println("dfs --- ", node.Val)
	fmt.Println("bfs", n.Val)

	node = lowestCommonAncestor(six, two, four)
	n = lca(six, two, four)
	fmt.Println("dfs --- ", node.Val)
	fmt.Println("bfs", n.Val)

	node = lowestCommonAncestor(six, zero, nine)
	n = lca(six, zero, nine)
	fmt.Println("dfs --- ", node.Val)
	fmt.Println("bfs", n.Val)

	node = lowestCommonAncestor(six, seven, nine)
	n = lca(six, seven, nine)
	fmt.Println("dfs --- ", node.Val)
	fmt.Println("bfs", n.Val)
}