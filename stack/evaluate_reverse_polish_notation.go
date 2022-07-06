package main

import (
	"fmt"
	"strconv"
)

func evalPRN(tokens []string) int {
	// Idea: add the numbers to the stack. How the input works is we will have 
	// two numbers in the stack before we get to operators ( + , - , * , / )
	// access both numbers and operate. Save the information in the stack on the second
	// to last item and then remove last item from the stack
	// return the last item into the stack
	stack := []int{}
	for _, token := range tokens {
		number, err := strconv.Atoi(token)
		if err != nil {
			l, r := stack[len(stack) - 2], stack[len(stack) - 1]
			switch token {
			case "+":
				stack[len(stack) - 2] = l + r
			case "-":
				stack[len(stack) - 2] = l - r
			case "*":
				stack[len(stack) - 2] = l * r
			case "/":
				stack[len(stack) - 2] = l / r
			}
			stack = stack[:len(stack) - 1]
		} else {
			stack = append(stack, number)
		}
	}
	return stack[len(stack) - 1]
}

func main() {
	
	t1 := []string{"2","1","+","3","*"}
	t2 := []string{"4","13","5","/","+"}
	t3 := []string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}

	fmt.Println(evalPRN(t1)) // 9
	fmt.Println(evalPRN(t2)) // 6
	fmt.Println(evalPRN(t3)) // 22
}