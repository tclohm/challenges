// Binary trees are already defined with this interface:
// type Tree struct {
//   Value interface{}
//   Left *Tree
//   Right *Tree
// }
type Stack struct {
	nodes []*Tree
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Tree) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Tree {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}


func kthSmallestInBST(t *Tree, k int) int {
    stack := NewStack()

    cur := t
    stack.Push(t)
    for stack.count > 0 || cur != nil {
        if (cur != nil) {
            stack.Push(cur)
            cur = cur.Left
        } else {
            cur = stack.Pop()
            k--
            
            if k == 0 {
                return cur.Value.(int)
            }
            
            cur = cur.Right
        }
    }
    
    return 0
}



func kthSmallestInBST(t *Tree, k int) int {
    n := 0
    for (t != nil){
        if t.Left == nil{
            n += 1
            if n == k{
            return t.Value.(int)
        }
            t = t.Right
        } else {
            child := t.Left
            for (child.Right != nil && child.Right != t) {

                child = child.Right
            }

            if (child.Right == nil){
                child.Right = t;
                t = t.Left;

            } else {
                n += 1
                child.Right = nil;
                if n == k{
                    return t.Value.(int)
                }
                t = t.Right;
            }

        }
        
    }
    return 0
}
