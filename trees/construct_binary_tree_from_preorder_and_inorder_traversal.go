package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func indexOf(target int, in []int) int {
	for i := 0 ; i < len(in) ; i++ {
		if target == in[i] {
			return i
		}
	}
	return -1
}
 
func buildTree(preorder []int, inorder []int) *TreeNode {
    // preorder root -> go left until no more, -> go right until no more

    // inorder left -> root -> right
    if len(preorder) == 0 || len(inorder) == 0 {
    	return nil
    }

    index := indexOf(preorder[0], inorder)
    // recursive
    node := &TreeNode{
    	Val: preorder[0],
    	Left: buildTree(preorder[1:index+1], inorder[:index]),
    	Right: buildTree(preorder[index+1:], inorder[index+1:]),
    }

    return node
}



func main() {
	preorder := []int{3,9,20,15,7}
	inorder := []int{9,3,15,20,7}

	buildTree(preorder, inorder)
}