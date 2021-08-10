// Singly-linked lists are already defined with this interface:
// type ListNode struct {
//   Value interface{}
//   Next *ListNode
// }
//
func removeKFromList(l *ListNode, k int) *ListNode {
   // * and & are pointers
   h := &ListNode{}
   h.Next = l
   l = h
   for l.Next != nil {
      if l.Next.Value == k {
         l.Next = l.Next.Next
      } else {
         l = l.Next
      }

   }
   return h.Next
}

func removeKFromList(l *ListNode, k int) *ListNode {
    if l == nil {
        return nil
    }
    if l.Value == k {
        return removeKFromList(l.Next, k)
    } else {
        l.Next = removeKFromList(l.Next, k)
        return l
    }
    
}