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

func isSame(p, q *TreeNode) bool {
	if p == nil && q == nil { return true }
	if p == nil && q != nil || p != nil && q == nil { return false }
	return p.Val == q.Val && isSame(p.Left, q.Left) && isSame(p.Right, q.Right)
}

func main() {
	th1 := &TreeNode{Val:3}
	t1 := &TreeNode{Val: 2}
	r1 := &TreeNode{Val: 1, Left: t1, Right: th1}

	th2 := &TreeNode{Val:3}
	t2 := &TreeNode{Val: 2}
	r2 := &TreeNode{Val: 1, Left: t2, Right: th2}
	fmt.Println(isSame(r1, r2))

	tw1 := &TreeNode{Val: 2}
	rt1 := &TreeNode{Val: 1, Left: tw1}
	tw2 := &TreeNode{Val: 2}
	rt2 := &TreeNode{Val: 1, Right: tw2}
	fmt.Println(isSame(rt1, rt2))
}