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
	return len(s.container) == 0
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func maxDepth(s string) int {
	depth := 0
	deepest := 0
	stack := &stack{}
    for _, r := range s {
    	if !stack.isEmpty() && r == ')'{
    		if stack.pop() == '(' {
    			depth--
    		}
    	}
    	if r == '(' {
    		stack.push(r)
    		depth++
    	}
    	deepest = max(deepest, depth)
    }

    return deepest
}

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
	fmt.Println(maxDepth("(1)+((2))+(((3)))"))
}