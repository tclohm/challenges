package main

import "fmt"

func containsOpen(s rune) bool {
	openSymbols := []rune{'{', '[', '('}
	for _, r := range openSymbols {
		if s == r {
			return true
		}
	}
	return false
}

func popAndCompare(s []rune, r rune) ([]rune, bool) {
	last := len(s) - 1
	openSymbol := s[last]
	s = s[:last]

	if openSymbol == '{' && r == '}' {
		return s, true
	}

	if openSymbol == '[' && r == ']' {
		return s, true
	}

	if openSymbol == '(' && r == ')' {
		return s, true
	}

	return s, false
}

func isValid(s string) bool {
	stack := []rune{}

	for _, symbol := range s {
		if containsOpen(symbol) {
			stack = append(stack, symbol)
		} else {
			mutatedStack, validity := popAndCompare(stack, symbol); if !validity {
				return false
			}
			stack = mutatedStack
		}
	}
	return len(stack) == 0
}

type stack []int32
func (s *stack) Push(v int32) {
	*s = append(*s, v)
}
func (s *stack) Pop() {
	if s.IsEmpty() {
		return
	}
	*s =  (*s)[len(*s) - 1]
}
func (s *stack) Top() int32 {
	if s.IsEmpty() {
		return 0
	}
	return (*s)[len(*s) - 1]
}
func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func validity(s string) bool {
	stck := make(stack, 0)
	for _, v := range s {
		switch {
		case v == '(':
			stck.Push(')')
		case v == '{':
			stck.Push('}')
		case v == '[':
			stck.Push(']')
		default:
			if stck.IsEmpty() || v != stck.Top() {
				return false
			} else {
				stck.Pop()
			}
		}
	}
	return stck.IsEmpty()
}

func main() {
	v := "()"
	vv := "()[]{}"
	nv := "(]"
	vvv := "[]((({})))()"

	fmt.Println(isValid(v))
	fmt.Println(isValid(vv))
	fmt.Println(isValid(nv))
	fmt.Println(isValid(vvv))
}