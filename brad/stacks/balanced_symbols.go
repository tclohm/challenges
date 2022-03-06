package main

import "fmt"

func is_balanced(symbols string) bool {
	stack := []rune{}

	pairings := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, s := range symbols {
		if _, ok := pairings[s]; ok {
			stack = append(stack, s)
			continue
		}

		if len(stack) > 0 {
			top := len(stack) - 1
			expected_opening_symbol := stack[top]
			stack = stack[:top]

			if s != pairings[expected_opening_symbol] {
				fmt.Println(string(s), string(pairings[expected_opening_symbol]))
				return false
			}

		} else {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println("{{([][])}()} is balanced:", is_balanced("{{([][])}()}"))
	fmt.Println("[[{{(())}}]] is balanced:", is_balanced("[[{{(())}}]]"))
	fmt.Println("[][][](){}   is balanced:", is_balanced("[][][](){}"))
	fmt.Println("---")
	fmt.Println("{[]) 	is balanced:", is_balanced("{[])"))
	fmt.Println("((()))	is balanced:", is_balanced("((()))"))
	fmt.Println("(() 	is balanced:", is_balanced("(()"))
	fmt.Println("()) 	is balanced:", is_balanced("())"))
}