package main

import "fmt"

type Stack struct {
	container []interface{}
}

func (s *Stack) print(msg string) {
	fmt.Println(msg, "\nstack:", s.container)
}

func (s *Stack) Pop() interface{} {
	end := len(s.container) - 1
	item := s.container[end]
	s.container = s.container[0:end]
	s.print("pop")
	return item
}

func (s *Stack) Push(item interface{}) {
	s.container = append(s.container, item)
	s.print("push")
} 

func isValid(s string) bool {
	pairs := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	st := Stack{}

	for _, r := range s {

		if opening, ok := pairs[r]; ok {
			fmt.Println(string(opening), "==", string(r))
			if opening != st.Pop() { return false }
		} else {
			st.Push(r)
		}
		
	}

	return true
}

func main() {
	not_valid := "{([])]"
	valid := "{([])}"
	fmt.Println("\tNOT VALID")
	fmt.Println(isValid(not_valid))
	fmt.Println("\n\tVALID")
	fmt.Println(isValid(valid))
}