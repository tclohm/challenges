package main

import "fmt"

func is_balanced(parentheses string) bool {
	stack := []rune{}

	OPENING := '('

	for _, paren := range parentheses {
		if paren == OPENING {
			stack = append(stack, paren)
		} else {
			if len(stack) > 0 {
				end := len(stack) - 1
				stack = stack[:end]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}


func main() {
	fmt.Println("((())) is balanced? =>", is_balanced("((()))"))
	fmt.Println("(()    is balanced? =>", is_balanced("(()"))
	fmt.Println("())    is balanced? =>", is_balanced("())"))
	fmt.Println("(()()()) is balanced? =>", is_balanced("(()()())"))
}