package main

import "fmt"

type BinaryNode struct {
	Value int
	Left *BinaryNode
	Right *BinaryNode
}

func (self *BinaryNode) preorderWalk(curr *BinaryNode, path *[]int) {
	
	if curr == nil { return }

	*path = append(*path, curr.Value)
	self.preorderWalk(curr.Left, path)
	self.preorderWalk(curr.Right, path)
}

func (self *BinaryNode) inorderWalk(curr *BinaryNode, path *[]int) {
	if curr == nil { return }
	self.inorderWalk(curr.Left, path)
	*path = append(*path, curr.Value)
	self.inorderWalk(curr.Right, path)
}

func (self *BinaryNode) postorderWalk(curr *BinaryNode, path *[]int) {
	if curr == nil { return }
	self.postorderWalk(curr.Left, path)
	self.postorderWalk(curr.Right, path)
	*path = append(*path, curr.Value)
}

func (self *BinaryNode) preorder() []int {
	var path = []int{}
	self.preorderWalk(self, &path)
	return path
}

func (self *BinaryNode) inorder() []int {
	var path = []int{}
	self.inorderWalk(self, &path)
	return path
}

func (self *BinaryNode) postorder() []int{
	var path = []int{}
	self.postorderWalk(self, &path)
	return path
}

func main() {
	seven := &BinaryNode{7, nil, nil}
	six := &BinaryNode{6, nil, nil}
	five := &BinaryNode{5, nil, nil}
	four := &BinaryNode{4, nil, nil}
	three := &BinaryNode{3, six, seven}
	two := &BinaryNode{2, four, five}
	one := &BinaryNode{1, two, three}
	fmt.Println("preorder")
	fmt.Println(one.preorder())
	fmt.Println("inorder")
	fmt.Println(one.inorder())
	fmt.Println("postorder")
	fmt.Println(one.postorder())
}