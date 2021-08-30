type Tree struct {
  Value interface{}
  Left *Tree
  Right *Tree
}


type Stack struct {
	nodes []*Tree
}

func NewStack {
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
	s = s[:len(s)-1]
	return node
}


func isLeaf(node *Tree) bool {
	if node != nil && node.Left == nil && node.Right == nil {
		return true
	}
	return false
}

func leafsum(root *Tree) int {

	if root == nil {
		return 0
	}

	stack := NewStack{}
	stack.Push(root)
	total := 0


	for len(stack) != 0 {
		node := stack.Pop()
		if isLeaf(node) {
			total += node.Value.(int)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
		if node.Right != nil {
			stack.Push(node.Right)
		}
	}

	return total
}
