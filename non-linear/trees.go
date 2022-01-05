package main

import (
	"fmt"
	"sync"
)

type Node struct {
	key int
	value int
	left *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

func (tree *BinarySearchTree) Insert(key int, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	var node *Node

	node = &Node{key, value, nil, nil}

	if tree.root == nil {
		tree.root = node
	} else {
		insertNode(tree.root, node)
	}
}

func insertNode(root *Node, newNode *Node) {
	if newNode.key < root.key {
		if root.left == nil {
			root.left = newNode
		} else {
			insertNode(root.left, newNode)
		}
	} else {
		if root.right == nil {
			root.right = newNode
		} else {
			insertNode(root.right, newNode)
		}
	}
}

func (tree *BinarySearchTree) InOrderTraverseTree(function func(int)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	inOrderTraverseTree(tree.root, function)
}

func inOrderTraverseTree(node *Node, function func(int)) {
	if node != nil {
		inOrderTraverseTree(node.left, function)
		function(node.value)
		inOrderTraverseTree(node.right, function)
	}
}

func (tree *BinarySearchTree) PreOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	preOrderTraverseTree(tree.root, function)
}

func preOrderTraverseTree(node *Node, function func(int)) {
	if node != nil {
		function(node.value)
		preOrderTraverseTree(node.left, function)
		preOrderTraverseTree(node.right, function)
	}
}

func (tree *BinarySearchTree) PostOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	postOrderTraverseTree(tree.root, function)
}

func postOrderTraverseTree(node *Node, function func(int)) {
	if node != nil {
		postOrderTraverseTree(node.left, function)
		postOrderTraverseTree(node.right, function)
		function(node.value)
	}
}

func (tree *BinarySearchTree) MinNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	
	var node *Node
	
	if node == nil {
		return (*int)(nil)
	}

	for {
		if node.left == nil {
			return &node.value
		}
		node = node.left
	}
}

func (tree *BinarySearchTree) MaxNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	var node *Node
	node = tree.root

	if node == nil {
		return (*int)(nil)
	}

	for {
		if node.right == nil {
			return &node.value
		}
		node = node.right
	}
}

func (tree *BinarySearchTree) search(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return search(tree.root, key)
}

func search(node *Node, key int) bool {
	if node == nil {
		return false
	}
	if key < node.key {
		return search(node.left, key)
	}
	if key > node.key {
		return search(node.right, key)
	}
	return true
}

func (tree *BinarySearchTree) Remove(key int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	remove(tree.root, key)
}

func remove(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = remove(node.left, key)
		return node
	}
	if key > node.key {
		node.right = remove(node.right, key)
		return node
	}
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}
	var leftMostRightNode *Node
	leftMostRightNode = node.right
	for {
		if leftMostRightNode != nil && leftMostRightNode.left != nil {
			leftMostRightNode = leftMostRightNode.left
		} else {
			break
		}
	}
	node.key, node.value = leftMostRightNode.key, leftMostRightNode.value
	node.right = remove(node.right, node.key)
	return node
}

func (tree *BinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Println("------------")
	stringify(tree.root, 0)
	fmt.Println("------------")
}

func stringify(node *Node, level int) {
	if node != nil {
		format := ""
		for i := 0 ; i < level ; i++ {
			format += " "
		}
		format += "----> "
		level++
		stringify(node.left, level)
		fmt.Printf(format+"%d\n", node.key)
		stringify(node.right, level)
	}
}

func main() {
	var tree *BinarySearchTree = &BinarySearchTree{}
	tree.Insert(8,8)
	tree.Insert(3,3)
	tree.Insert(10,10)
	tree.Insert(1,1)
	tree.Insert(6,6)
	tree.String()
}