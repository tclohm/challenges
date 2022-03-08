package main

import (
	"container/list"
	"fmt"
)

func is_palindrome(characters string) bool {
	deque := list.New()

	for _, char := range characters {
		deque.PushBack(string(char))
	}

	for i, j := deque.Front(), deque.Back(); i != deque.Back() ; i, j = i.Next(), j.Prev() {
		if i.Value != j.Value {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(is_palindrome("radar"))
	fmt.Println(is_palindrome("lsdkjfskf"))
}