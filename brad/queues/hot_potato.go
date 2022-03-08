package main

import (
	"container/list"
	"fmt"
)

func hot_potato(names []string, number int) string {
	queue := list.New()

	for _, name := range names {
		queue.PushFront(name)
	}


	for queue.Len() > 1 {
		for i := 0 ; i < number ; i++ {
			queue.PushFront(queue.Remove(queue.Back()))
		}
		queue.Remove(queue.Back())
	}

	name :=  fmt.Sprintf("%v", queue.Back().Value)

	return name
}

func main() {
	names := []string{"Bill", "David", "Susan", "Jane", "Kent", "Brad"}
	fmt.Println(hot_potato(names, 9))
}