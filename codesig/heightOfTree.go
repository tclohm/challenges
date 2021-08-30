type Tree struct {
	Value interface{}
	Left *Tree
	Right *Tree
}

type Stack struct {
	nodes []*Tree
}

func Stack {
	return &Stack{}
}

func (s *Stack) Push(node *Tree) {
	s.nodes = append(s.nodes, node)
}

func (s *Stack) Pop() *Tree {
	if len(s) == 0 {
		return nil
	}

	node := s[len(s) - 1]
	s = s[:len(s) - 1]
	return node
}

func isLeaf(node *Tree) bool {
	if node != nil && node.Left == nil && node.Right == nil {
		return true
	}
	return false
}

func treeHeight(root *Tree) int {

	if isLeaf(root){
		return 0
	}

	return math.Max(treeHeight(node.left), treeHeight(node.Right)) + 1
}

func treeHeightUsingStack(root *Tree) int {
	stack := NewStack{root}
	depth := 0
	tmpDepth := 0

	for len(stack) > 0 {

		node := stack.Pop()
		tmpDepth += 1

		if isLeaf(node) {
			depth = math.Max(tmpDepth, depth)
			tmpDepth = 0
		}

		if node.Left != nil {
			stack.Push(node.Left)
		}

		if node.Right != nil {
			stack.Push(node.Right)
		}

	}
	return depth
}