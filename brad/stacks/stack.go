package main

import "fmt"

type Stack struct {
	_items []interface{}
}

func (s *Stack) IsEmpty() bool {
	return len(s._items) == 0
}

func (s *Stack) Push(item interface{}) {
	fmt.Println("pushing", item, "on to", s)
	s._items = append(s._items, item)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return nil
	}
	removed := s._items[len(s._items) - 1]
	fmt.Println("Popping", removed)
	s._items = s._items[:len(s._items) - 1]
	return removed
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return nil
	}

	return s._items[len(s._items) - 1]
}

func (s *Stack) Size() int {
	return len(s._items)
}

func main() {
	s := Stack{}
	s.Push(1)
	s.Push(40)
	s.Pop()
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Peek())
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	fmt.Println(s.Size())
}