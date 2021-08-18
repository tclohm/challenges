//
// Binary trees are already defined with this interface:
// type Tree struct {
//   Value interface{}
//   Left *Tree
//   Right *Tree
// }
func hasPathWithGivenSum(t *Tree, s int) bool {
    
    if t == nil {
        return false
    }
    
    path := []int{}
    queue := []*Tree{t}
    
    for len(queue) != 0 {
        node := queue[0]
        queue = queue[1:]
        
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
        
        if node.Value != nil {
            path = append(path, node.Value.(int))
        }
        
    }
    
    total := 0
    
    for _, v := range path {
        total += v
    }
    
    if total == s {
        return true
    }
    
    return false
}


// Actual stack
type TreeStack struct {
    t *Tree
    sumAncestor int
}

type stack []*TreeStack

func (s stack) Push(v *TreeStack) stack {
    return append(s, v)
}

func (s stack) Pop() (stack, *TreeStack) {
    l := len(s)
    if l == 0 {
        return nil, nil
    }
    return s[:l-1], s[l-1]
}

func (s stack) IsEmpty() bool {
    if s == nil {
        return true
    }
    return len(s) == 0
}

func hasPathWithGivenSum(t *Tree, s int) bool {
    // precondition
    if (t == nil) {
        return false
    }
    // initialization
    myStack := stack{}.Push(&TreeStack{t, 0})
    // until the stack is empty
    for !myStack.IsEmpty() {
        newStack, tmp := myStack.Pop()
        myStack = newStack
        sum := tmp.sumAncestor + tmp.t.Value.(int)
        // we push to the stack if possible
        if tmp.t.Right != nil {
            myStack = myStack.Push(&TreeStack{tmp.t.Right, sum})
        }
        if tmp.t.Left != nil {
            myStack = myStack.Push(&TreeStack{tmp.t.Left, sum})
        }
        if (tmp.t.Left == nil && tmp.t.Right == nil) {
            // edge of a graph
            if sum == s {
                fmt.Println("should break")
                return true
            }
        }
        fmt.Println("end", sum)
    }
    fmt.Println("finished")
    return false
}