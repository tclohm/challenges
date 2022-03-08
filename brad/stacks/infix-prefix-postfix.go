package main

import (
	"fmt"
	"strings"
)

func exist(target, search string) bool {
	for _, char := range search {
		if string(char) == target {
			return true
		}
	}
	return false
}

func infix_to_postfix(infix_expression string) string {
	PRECEDENCE := map[string]int{
		"*": 3,
		"/": 3,
		"+": 2,
		"-": 2,
		"(": 1,
	}

	CHARACTERS := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DIGITS := "0123456789"
	LEFT_PAREN := "("
	RIGHT_PAREN := ")"

	operation_stack := []string{}
	postfix := []string{}
	tokens := strings.Split(infix_expression, "")

	for _, token := range tokens {
		string_token := string(token)
		if exist(string_token, CHARACTERS) || exist(string_token, DIGITS) {
			postfix = append(postfix, string_token)
		} else if string_token == LEFT_PAREN {
			operation_stack = append(operation_stack, string_token)
		} else if string_token == RIGHT_PAREN {
			top := len(operation_stack) - 1
			top_token := operation_stack[top]
			operation_stack = operation_stack[:top]

			for top_token != LEFT_PAREN {
				postfix = append(postfix, top_token)
				// pop
				top = len(operation_stack) - 1
				top_token = operation_stack[top]
				operation_stack = operation_stack[:top]
			}
		} else {
			top := len(operation_stack) - 1
			for len(operation_stack) > 0 && (PRECEDENCE[operation_stack[top]] >= PRECEDENCE[string_token]) {
				element := operation_stack[top]
				operation_stack = operation_stack[:top]
				top = len(operation_stack) - 1
				postfix = append(postfix, element)
			}
			operation_stack = append(operation_stack, string_token)
		}
	}

	for len(operation_stack) > 0 {
		top := len(operation_stack) - 1
		postfix = append(postfix, operation_stack[top])
		operation_stack = operation_stack[:top]
	}

	return strings.Join(postfix, "")
}

func main() {
	fmt.Println(infix_to_postfix("A * B + C * D"))
	//fmt.Println(infix_to_postfix("( A + B ) * C"))
	fmt.Println(infix_to_postfix("A + B * C"))
}