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