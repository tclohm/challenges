// Singly-linked lists are already defined with this interface:
// type ListNode struct {
//   Value interface{}
//   Next *ListNode
// }
//
func reverseNodesInKGroups(l *ListNode, k int) *ListNode {
    // reverse l k at a time, k = 2 [1,2,3,4,5,6,7] => [2,1,4,3,6,5,7]
    length := length(l)
    nodesNotToBeFlipped := length % k
    threshold := length - nodesNotToBeFlipped
    pointer := l
    
    return reverse(pointer, k, 0, threshold)
}

func length(l *ListNode) int {
    count := 0
    
    for current := l ; current != nil ; current = current.Next {
        count++
    }
    
    return count
}

func reverse(pointer *ListNode, k, index, threshold int) *ListNode {
    if index + k > threshold {
        return pointer
    }
    
    current := pointer
    count := 0
    var previous *ListNode
    
    for current != nil && count < k {
        // reverse
        temporary := current.Next
        current.Next = previous
        previous = current
        current = temporary
        count++
        index++
    }
    
    if current != nil {
        pointer.Next = reverse(current, k, index, threshold)
    }
    
    return previous
}
