package main

import "fmt"

type stack struct {
	container []rune
}

func (s *stack) pop() rune {
	end := len(s.container) - 1
	el := s.container[end]
   	s.container = s.container[:end]
   	return el
}

func (s *stack) push(v rune) {
	s.container = append(s.container, v)
}

func (s *stack) isEmpty() bool {
	return s.length() == 0
}

func (s *stack) length() int {
	return len(s.container)
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func maxDepth(s string) int {
	deepest := 0
	stack := &stack{}
    for _, r := range s {
    	if r == ')' {
    		deepest = max(deepest, stack.length())
    		stack.pop()
    	}
    	if r == '(' {
    		stack.push(r)
    	}
    }

    return deepest
}

func noStackSolution(s string) int {
	deepest := 0
	currentDepth := 0

	for _, el := range s {
		if el == '(' {
			currentDepth++
		} else if el == ')' {
			deepest = max(currentDepth, deepest)
			currentDepth--
		}
	}

	return deepest
}

func main() {
	fmt.Println(noStackSolution("(1+(2*3)+((8)/4))+1"))
	fmt.Println(noStackSolution("(1)+((2))+(((3)))"))
}