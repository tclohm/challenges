package main

import "fmt"

func generate(n int) []string {
	answer := make([]string, 0)
	stack := make([]byte, 0)
	current := ""
	var backtrack func()
	backtrack = func() {

		if len(current) == n * 2 {
			if len(stack) == 0 {
				answer = append(answer, current)
			}
			return
		}

		if len(stack) > 0 && stack[len(stack) - 1] == '(' {
			current += ")"
			stack = stack[:len(stack) - 1]
			backtrack()
			current = current[:len(current) - 1]
			stack = append(stack, '(')
		}
		current += "("
		stack = append(stack, '(')
		backtrack()
		stack = stack[:len(stack) - 1]
		current = current[:len(current) - 1]
	}
	backtrack()
	return answer
}

func main() {
	fmt.Println(generate(3))
	fmt.Println(generate(1))
}