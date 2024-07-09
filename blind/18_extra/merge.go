package main

import (
	"fmt"
	"bytes"
)

type Node struct {
	val int
	next *Node
}

func printLinkedList(head *Node) {
	var buffer bytes.Buffer
	for head != nil {
		buffer.WriteString(fmt.Sprintf("%d -> ", head.val))
		head = head.next
	}
	buffer.WriteString("nil")

	fmt.Println(buffer.String())
}

func main() {
	four := Node{4, nil}
	three := Node{3, &four}
	two := Node{2, &three}
	one := Node{1, &two}
	printLinkedList(&one)
}
