package main

type Value interface {
	LessThan(Value) bool
	EqualTo(Value) bool
}

type Node struct {
	Key Value
	Balance Value
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
	var opposite *Node
	opposite = node.LinkedNodes[opposite(balanceValue)]
	switch opposite.BalanceValue {
	case 0:
		root.BalanceValue = 0
		node.BalanceValue = 0
	case balanceValue:
		root.BalanceValue = -balanceValue
		node.BalanceValue = 0
	default:
		root.BalanceValue = 0
		node.BalanceValue = balanceValue
	}
	opposite.BalanceValue = 0
}