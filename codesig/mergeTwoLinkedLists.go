// Singly-linked lists are already defined with this interface:
// type ListNode struct {
//   Value interface{}
//   Next *ListNode
// }
//
func exhaust(linkedlist **ListNode, into **ListNode) {
    for *linkedlist != nil {
        (*into).Value = (*linkedlist).Value
        
        if (*linkedlist).Next != nil {
            (*into).Next = &ListNode{}
        }

        *linkedlist = (*linkedlist).Next
        *into = (*into).Next      
    }
    return
}

func mergeAndNext(linkedlist **ListNode, into **ListNode) {
    (*into).Value = (*linkedlist).Value
    *into = (*into).Next
    *linkedlist = (*linkedlist).Next
    return
}

func mergeTwoLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{}
    head := result
    if l1 == nil && l2 == nil {
        return l1
    }
    for l2 != nil || l1 != nil {
        if l1 != nil && l2 != nil {
            if l1.Value.(int) <= l2.Value.(int) {
                result.Next = &ListNode{}
                mergeAndNext(&l1, &result)
            } else {
                result.Next = &ListNode{}
                mergeAndNext(&l2, &result)
            }
        }
        
        if l1 == nil && l2 != nil {
            exhaust(&l2, &result)
        } 
        
        if l1 != nil && l2 == nil {
            exhaust(&l1, &result)
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


func mergeTwoLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{}
    r := res
    for l1 != nil || l2 != nil {
        r.Next = &ListNode{min(&l1, &l2), nil}
        r = r.Next
    }
    return res.Next
}

func min(l1 **ListNode, l2 **ListNode) (minVal int) {
    if *l1 == nil || (*l2 != nil && (*l1).Value.(int) > (*l2).Value.(int)) {
        minVal = (*l2).Value.(int)
        *l2 = (*l2).Next
        return
    } else {
        minVal = (*l1).Value.(int)
        *l1 = (*l1).Next
        return
    }
}