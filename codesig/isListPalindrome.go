// Singly-linked lists are already defined with this interface:
// type ListNode struct {
//   Value interface{}
//   Next *ListNode
// }
//
func isListPalindrome(l *ListNode) bool {
   nums := []int{}
   for l != nil {
       n, _ := l.Value.(int)
       nums = append(nums, n)
       l = l.Next
   }
   
   for i := 0 ; i < len(nums) ; i++ {
       if nums[i] != nums[len(nums) - 1 - i] {
           return false
       }
   }
   
   return true
}

func isListPalindrome(l *ListNode) bool {
    fast := l
    reversed := &ListNode{}
    temp := l
    for fast != nil && fast.Next != nil {
        // move forward twice
        fast = fast.Next.Next
        // temporarily grab the next node
        temp = l.Next
        // list next will point to reversed pointer, first one is nil
        l.Next = reversed
        // slow becomes l
        reversed = l
        // l becomes temp
        l = temp
    }
    
    if fast != nil {
        l = l.Next
    }
    
    for curr := l ; curr != nil ; curr = curr.Next {
        if reversed.Value != curr.Value {
            return false
        }
        reversed = reversed.Next
    }
    
    return true
}
