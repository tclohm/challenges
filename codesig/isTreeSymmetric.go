//
// Binary trees are already defined with this interface:
// type Tree struct {
//   Value interface{}
//   Left *Tree
//   Right *Tree
// }
func isTreeSymmetric(t *Tree) bool {
    
    if t == nil {
        return true
    }
    
    leftside := []*Tree{}
    rightside := []*Tree{}
    
    if t != nil && t.Left != nil {
        leftside = append(leftside, t.Left)
    }
    
    if t != nil && t.Right != nil {
        rightside = append(rightside, t.Right)
    }
    
    for len(leftside) != 0 || len(rightside) != 0 {
        
        if len(leftside) != len(rightside) {
            return false
        }
        
        left := leftside[0]
        right := rightside[0]
        
        leftside = leftside[1:]
        rightside = rightside[1:]
        
        if left.Value != right.Value {
            return false
        }
        
        if left.Left == nil && right.Right != nil {
            return false
        }
        
        if left.Right == nil && right.Left != nil {
            return false
        }
        
        if ( left != nil && left.Left != nil ) && ( right != nil && right.Right != nil ) {
            leftside = append(leftside, left.Left)
            rightside = append(rightside, right.Right)
        }
        
        if ( left != nil && left.Right != nil ) && ( right != nil && right.Left != nil ) {
            leftside = append(leftside, left.Right)
            rightside = append(rightside, right.Left)
        }
    }
    
    return true
}


func isTreeSymmetric(t *Tree) bool {
    return helper(t, t)
}

func helper (left *Tree, right *Tree) bool {
    if left == nil && right == nil {
        return true
    }
    
    if left == nil || right == nil || left.Value != right.Value {
        return false
    }
    
    return helper(left.Left, right.Right) && helper(left.Right, right.Left)
}
