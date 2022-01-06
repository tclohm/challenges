package main

import (
	"fmt"
	"encoding/json"
)

type Value interface {
	LessThan(Value) bool
	EqualTo(Value) bool
}

type Node struct {
	Key Value
	Balance int
	LinkedNodes [2]*Node
}
// takes node value and returns opposite
func opposite(value int) int {
	return 1 - value
}
// rotate node opposite to the specified sub-tree
func singleRotation(root *Node, value int) *Node {
	var saveNode *Node
	saveNode = root.LinkedNodes[opposite(value)]
	root.LinkedNodes[opposite(value)] = saveNode.LinkedNodes[value]
	saveNode.LinkedNodes[value] = root
	return saveNode
}

func doubleRotation(root *Node, value int) *Node {
	var saveNode *Node
	saveNode = root.LinkedNodes[opposite(value)].LinkedNodes[value]
	root.LinkedNodes[opposite(value)].LinkedNodes[value] = root.LinkedNodes[opposite(value)]
	saveNode.LinkedNodes[opposite(value)] = root.LinkedNodes[opposite(value)]
	root.LinkedNodes[opposite(value)] = saveNode
	saveNode = root.LinkedNodes[opposite(value)]
	root.LinkedNodes[opposite(value)] = saveNode.LinkedNodes[value]
	saveNode.LinkedNodes[value] = root
	return saveNode
}

func adjust(root *Node, value int, balanceValue int) {
	var node *Node
	node = root.LinkedNodes[value]
	var opp *Node
	opp = node.LinkedNodes[opposite(balanceValue)]
	switch opp.Balance {
	case 0:
		root.Balance = 0
		node.Balance = 0
	case balanceValue:
		root.Balance = -balanceValue
		node.Balance = 0
	default:
		root.Balance = 0
		node.Balance = balanceValue
	}
	opp.Balance = 0
}

func BalanceTree(root *Node, value int) *Node {
	var node *Node

	node = root.LinkedNodes[value]

	var balance int

	balance = 2 * value - 1
	
	if node.Balance == balance {
		root.Balance = 0
		node.Balance = 0
		return singleRotation(root, opposite(value))
	}
	adjust(root, value, balance)
	return doubleRotation(root, opposite(value))
}

func insertRight(root *Node, key Value) (*Node, bool) {
	if root == nil {
		return &Node{Key: key}, false
	}

	var dir int
	dir = 0

	if root.Key.LessThan(key) {
		dir = 1
	}

	var done bool

	root.LinkedNodes[dir], done = insertRight(root.LinkedNodes[dir], key)

	if done {
		return root, true
	}

	root.Balance = root.Balance + (2 * dir - 1)

	switch root.Balance {
	case 0:
		return root, true
	case 1, -1:
		return root, false
	}

	return BalanceTree(root, dir), true
}

func Insert(node **Node, key Value) {
	*node, _ = insertRight(*node, key)
}

func Remove(node **Node, key Value) {
	*node, _ = removeR(*node, key)
}

func removeBalance(root *Node, value int) (*Node, bool) {
	var node *Node
	node = root.LinkedNodes[opposite(value)]

	var balance int
	balance = 2 * value - 1

	switch node.Balance {
	case -balance:
		root.Balance = 0
		node.Balance = 0
		return singleRotation(root, value), false
	case balance:
		adjust(root, opposite(value), -balance)
		return doubleRotation(root, value), false
	}

	root.Balance = -balance
	node.Balance = balance

	return singleRotation(root, value), true
}

func removeR(root *Node, key Value) (*Node, bool) {
	if root == nil {
		return nil, false
	}

	if root.Key.EqualTo(key) {
		switch {
		case root.LinkedNodes[0] == nil:
			return root.LinkedNodes[1], false
		case root.LinkedNodes[1] == nil:
			return root.LinkedNodes[0], false
		}

		var heir *Node
		heir = root.LinkedNodes[0]
		for heir.LinkedNodes[1] != nil {
			heir = heir.LinkedNodes[1]
		}

		root.Key = heir.Key

		key = heir.Key
	}

	var dir int
	dir = 0

	if root.Key.LessThan(key) {
		dir = 1
	}

	var done bool

	root.LinkedNodes[dir], done = removeR(root.LinkedNodes[dir], key)

	if done {
		return root, true
	}

	root.Balance = root.Balance + (1 - 2 * dir)

	switch root.Balance {
	case 1, -1:
		return root, true
	case 0:
		return root, false
	}

	return removeBalance(root, dir)
}

type integerKey int

func (k integerKey) LessThan(k1 Value) bool { return k < k1.(integerKey) }
func (k integerKey) EqualTo(k1 Value) bool { return k == k1.(integerKey) }


func main() {
	var tree *Node

	fmt.Println("Tree is empty")

	var avl []byte

	avl, _ = json.MarshalIndent(tree, "", " ")

	fmt.Println(string(avl))

	fmt.Println("\n Add Tree")

	Insert(&tree, integerKey(5))
	Insert(&tree, integerKey(3))
	Insert(&tree, integerKey(8))
	Insert(&tree, integerKey(7))
	Insert(&tree, integerKey(6))
	Insert(&tree, integerKey(10))

	avl, _ = json.MarshalIndent(tree, "", " ")
	fmt.Println(string(avl))

	fmt.Println("\n Delete Tree")

	Remove(&tree, integerKey(3))
	Remove(&tree, integerKey(7))

	avl, _ = json.MarshalIndent(tree, "", " ")
	fmt.Println(string(avl))
}