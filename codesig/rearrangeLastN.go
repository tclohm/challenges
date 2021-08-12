// Singly-linked lists are already defined with this interface:
// type ListNode struct {
//   Value interface{}
//   Next *ListNode
// }
//
func rearrangeLastN(l *ListNode, n int) *ListNode {
    // if l == nil or n == 0 return l
    if l == nil || n == 0 {
        return l
    }
    
    fast := l
    slow := l
    
    // move fast forward until we reach 0 with n
    for n > 0 && fast != nil {
        fast = fast.Next;
        n--;
    }
    
    // if n is still not 0 and fast is nil we are done
    if n >= 0 && fast == nil {
        return l
    }
    
    // start slow now
    for fast.Next != nil {
        slow = slow.Next;
        fast = fast.Next;
    }
    
    // our new head is slow.next
    head := slow.Next;
    // release slow.next from it's node
    slow.Next = nil;
    // reassign fast.Next to be the l
    fast.Next = l;
    
    return head;
}