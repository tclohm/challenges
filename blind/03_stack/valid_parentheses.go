package main

import "fmt"

type stack struct {
	container []string
}

func isOpen(c string) bool {
	open := map[string]bool{"(":true, "{":true, "[":true}
	if _, ok := open[c]; ok {
		return true
	}
	return false
}

func match(c string) string {
	closed := map[string]string{"(":")", "{":"}", "[":"]"}
	return closed[c]
}

func (s *stack) push(c string) {
	s.container = append(s.container, c)
}

func (s *stack) pop() string {
	if len(s.container) == 0 { return "" }
	last := len(s.container) - 1
	el := s.container[last]
	s.container = s.container[0:last]
	return el
}

func isValid(s string) bool {
	st := stack{}
	for _, r := range s {
		symbol := string(r)
		if isOpen(symbol) {
			st.push(symbol)
		} else {
			open := st.pop()
			if symbol != match(open) { return false }
		}
	}

	return true
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
}