package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

// Recursive way
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil && list2 != nil { return list2 }
	if list1 != nil && list2 == nil { return list1 }
	if list1 == nil && list2 == nil { return nil }
	newNode := &ListNode{}
	if list1.Val >= list2.Val {
		newNode.Val = list2.Val
		list2 = list2.Next
		newNode.Next = mergeTwoLists(list1, list2)
	} else {
		newNode.Val = list1.Val
		list1 = list1.Next
		newNode.Next = mergeTwoLists(list1, list2)
	}

	return newNode
}

func merge(l1, l2 *ListNode) *ListNode {
	var head *ListNode = &ListNode{} 
    var tail *ListNode = head
    
    for l1 != nil && l2 != nil {

        if l1.Val > l2.Val {
            tail.Next = l2
            l2 = l2.Next
        } else {
            tail.Next = l1
            l1 = l1.Next
        }
     
        tail = tail.Next
    }

   	if l1 != nil {
   		tail.Next = l1
   		tail = tail.Next
   	}

   	if l2 != nil {
   		tail.Next = l2
   		tail = tail.Next
   	}

	return head.Next
}


func __print__(current *ListNode) string {
	result := ""
	for current != nil {
		if current.Next != nil {
			result +=  fmt.Sprintf("%v -> ", current.Val)
		} else {
			result += fmt.Sprintf("%v", current.Val)
		}
		
		current = current.Next
	}
	return result
}

func main() {
	four := &ListNode{Val: 4, Next: nil}
	otherFour := &ListNode{Val: 4, Next: nil}
	three := &ListNode{Val: 3, Next: four}
	two := &ListNode{Val: 2, Next: otherFour}
	one := &ListNode{Val: 1, Next: two}
	oneToThree := &ListNode{Val: 1, Next: three}

	fmt.Println("recursive")
	recursiveMerged := mergeTwoLists(one, oneToThree)
	fmt.Println(__print__(recursiveMerged))
	fmt.Println("not recursive")
	merged := merge(one, oneToThree)
	fmt.Println(__print__(merged))


}