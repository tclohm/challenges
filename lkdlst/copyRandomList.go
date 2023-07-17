package main

func copyRandomList(head *Node) *Node {
	oldToCopy := make(map[*Node]*Node)

	current := head

	for current != nil {
		clone := &Node{Val: current.Val}
		oldToCopy[current] = clone
		current = current.Next
	}

	current = head

	for current != nil {
		clone := oldToCopy[current]
		clone.Next = oldToCopy[current.Next]
		clone.Random = oldToCopy[current.Random]
		current = current.Next
	}

	return oldToCopy[head]
}