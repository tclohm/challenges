package main

import (
	"fmt"
	"strings"
)

func makeGood(s string) string {
	stack := []string{}

	for curr, _ := range s {
		length := len(stack)
		last := len(stack) - 1
		if length > 0 && stack[last] != string(s[curr]) && strings.ToLower(stack[last]) == strings.ToLower(string(s[curr])) {
			stack = stack[:last]
		} else {
			stack = append(stack, string(s[curr]))
		}
	}

	return strings.Join(stack, "")
}

func main() {
	fmt.Println(makeGood("leEeetcode"))
	fmt.Println(makeGood("abBAcC"))
	fmt.Println(makeGood("s"))
}