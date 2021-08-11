func addTwoHugeNumbers(a *ListNode, b *ListNode) *ListNode {
    a = reverse(a)
    b = reverse(b)
    
    var currentA = a
    var currentB = b
    var result = &ListNode{}
    var currentResult = result
    var carrier int = 0
    
    for currentA != nil && currentB != nil {
        var sum = currentA.Value.(int) + currentB.Value.(int) + carrier
        if sum > 9999 {
            carrier = 1
            sum = getLastFourDigits(sum)
        } else {
            carrier = 0
        }
        currentResult.Next = &ListNode{
            Value: sum,
        }
        currentResult = currentResult.Next
        currentA = currentA.Next
        currentB = currentB.Next
    }
    
    for currentA != nil {
        var sum = currentA.Value.(int) + carrier
        if sum > 9999 {
            carrier = 1
            sum = getLastFourDigits(sum)
        } else {
            carrier = 0
        }
        currentResult.Next = &ListNode{
            Value: sum,
        }
        currentResult = currentResult.Next
        currentA = currentA.Next
    }
    
    for currentB != nil {
        var sum = currentB.Value.(int) + carrier
        if sum > 9999 {
            carrier = 1
            sum = getLastFourDigits(sum)
        } else {
            carrier = 0
        }
        currentResult.Next = &ListNode{
            Value: sum,
        }
        currentResult = currentResult.Next
        currentB = currentB.Next
    }
    
    if carrier == 1 {
        currentResult.Next = &ListNode {
            Value: 1,
        }
    }
    
    r = r.Next
    r = reverse(r)
    
    // reconstruct the 2 arrays
    a = reverse(a)
    b = reverse(b)
    return r
}

func reverse(l *ListNode) *ListNode {
    var current = l
    var prev *ListNode
    for current != nil {
        var next = current.Next
        current.Next = prev
        prev = current
        current = next
    }
    return prev
}

func print(l *ListNode) {
    var current = l
    for current != nil {
        fmt.Println(current.Value.(int))
        current = current.Next
    }
}

func getLastFourDigits(n int) int {
    return n % 10000
}