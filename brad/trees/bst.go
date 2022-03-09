package main

type TreeNode struct {
	key int
	value interface{}
	left *TreeNode
	right *TreeNode
	parent *TreeNode
}

func (tn *TreeNode) IsLeftChild() bool {
	return tn.parent && tn.parent.left == tn
}

func (tn *TreeNode) IsRightChild() bool {
	return tn.parent && tn.parent.right == tn
}

func (tn *TreeNode) IsLeaf() bool {
	return !tn.right || !tn.left
}

func (tn *TreeNode) HasAnyChildren() bool {
	return tn.right || tn.left
}

func (tn *TreeNode) HasBothChildren() bool {
	return tn.left && tn.right
}

func (tn *TreeNode) HasOneChild() bool {
	return tn.HasAnyChildren() && !tn.HasBothChildren()
}

func (tn *TreeNode) ReplaceNodeData(key, val, left, right) {
	tn.key = key
	tn.val = val
	tn.left = left
	tn.right = right

	if tn.left {
		tn.left.parent = tn
	}

	if tn.right {
		tn.left.parent = tn
	}
}

// More to do