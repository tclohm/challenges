package main

import (
	"container/list"
	"fmt"
)

func main() {
	var linked *list.List
	linked = list.New()

	var element *list.Element = linked.PushBack(14)

	var front *list.Element = linked.PushFront(1)

	linked.InsertBefore(6, element)
	linked.InsertAfter(5, front)

	for curr := linked.Front() ; curr != nil ; curr = curr.Next() {
		fmt.Println(curr.Value)
	}
}