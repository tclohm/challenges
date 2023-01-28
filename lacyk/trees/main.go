package main

import "fmt"

type BinaryNode struct {
	Value int
	Left *BinaryNode
	Right *BinaryNode
}

type Queue struct {
	elements []*BinaryNode
}

func (self *Queue) length() int {
	return len(self.elements)
}

func (self *Queue) push(element *BinaryNode) {
	self.elements = append(self.elements, element)
}

func (self *Queue) shift() *BinaryNode {
	var element = self.elements[0]
	self.elements = self.elements[1:]
	return element
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

func (self *BinaryNode) bfs(item int) bool {
	return self.levels(self, item)
}

func (self *BinaryNode) levels(head *BinaryNode, needle int) bool{
	var q = Queue{}
	q.push(head)

	for q.length() > 0 {
		var curr *BinaryNode = q.shift()

		if curr.Value == needle {
			return true
		}

		if curr.Left != nil {
			q.push(curr.Left)
		}

		if curr.Right != nil {
			q.push(curr.Right)
		}
	}

	return false
}

func (self *BinaryNode) compare(a, b *BinaryNode) bool {
	if a == nil && b == nil { return true }
	if a == nil || b == nil { return false }
	if a.Value != b.Value { return false }

	return self.compare(a.Left, b.Left) && self.compare(a.Right, b.Right)
}

func (self *BinaryNode) search(curr *BinaryNode, needle int) bool {
	if curr == nil {
		return false
	}

	if curr.Value == needle {
		return true
	}

	if curr.Value < needle {
		return self.search(curr.Right, needle)
	}

	return self.search(curr.Left, needle)
}

func (self *BinaryNode) dfs(needle int) bool {
	return self.search(self, needle)
}

func main() {
	one := &BinaryNode{1, nil, nil}
	two := &BinaryNode{2, one, nil}
	four := &BinaryNode{4, nil, nil}
	five := &BinaryNode{5, four, nil}
	three := &BinaryNode{3, two, five}
	ten := &BinaryNode{10, nil, nil}
	seven := &BinaryNode{7, nil, nil}
	eight := &BinaryNode{8, seven, nil}
	nine := &BinaryNode{9, eight, ten}
	six := &BinaryNode{6, three, nine}
	fmt.Println("preorder")
	fmt.Println(six.preorder())
	fmt.Println("inorder")
	fmt.Println(six.inorder())
	fmt.Println("postorder")
	fmt.Println(six.postorder())

	fmt.Println(six.bfs(2))

	fmt.Println(six.dfs(4))
}