package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
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

func mergeKSortedList(lists []*ListNode) *ListNode {
	k := len(lists)
	h := new(minHeap)

	for i := 0 ; i < k ; i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}

	dummy := new(ListNode)
	current := dummy

	for h.Len() > 0 {
		tmp := heap.Pop(h).(*ListNode)
		if tmp.Next != nil {
			heap.Push(h, tmp.Next)
		}
		current.Next = tmp
		current = current.Next
	}

	return dummy.Next
}

type minHeap []*ListNode

func (h minHeap) Len() int { return len(h) }
func (h *minHeap) Push(x interface{}) {
	// use pointer receivers because they modify the slice's length, not just its content
	*h = append(*h, x.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}
func main() {

}