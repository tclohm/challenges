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
