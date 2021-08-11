func mergeTwoLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{}
    head := result
    
    if l1 == nil && l2 == nil {
        return l1
    }
    
    for l2 != nil || l1 != nil {
        
        if l1 != nil || l2 != nil {
            result.Next = &ListNode{}
        }
        
        if l1 != nil && l2 != nil {
            if l1.Value.(int) <= l2.Value.(int) {
                result.Value = l1.Value
                result = result.Next
                l1 = l1.Next
            } else {
                result.Value = l2.Value
                result = result.Next
                l2 = l2.Next
            }
        }
        
        if l1 == nil && l2 != nil {
            for l2 != nil {
                
                result.Value = l2.Value
                
                if l2.Next != nil {
                    result.Next = &ListNode{}
                }

                l2 = l2.Next
                result = result.Next
                
            }
        } 
        
        if l1 != nil && l2 == nil {
            for l1 != nil {
                
                result.Value = l1.Value
                
                if l1.Next != nil {
                    result.Next = &ListNode{}
                }

                l1 = l1.Next
                result = result.Next
                
            }
        }
    }
    
    return head
}


func mergeTwoLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    
    if l2 == nil {
        return l1
    }
    
    if l1.Value.(int) <= l2.Value.(int) {
        l1.Next = mergeTwoLinkedLists(l1.Next, l2)
        return l1
    } else {
        l2.Next = mergeTwoLinkedLists(l1, l2.Next)
        return l2
    }
}