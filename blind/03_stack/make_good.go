package main

import (
	"fmt"
	"strings"
)

func makeGood(s string) string {
	stack := []string{}

	for curr, _ := range s {
		length := len(stack)
		if length > 0 && stack[length - 1] != string(s[curr]) && strings.ToLower(stack[length - 1]) == strings.ToLower(string(s[curr])) {
			stack = stack[:length - 1]
		} else {
			stack = append(stack, string(s[curr]))
		}
	}

	return strings.Join(stack[:], "")
}

func main() {
	fmt.Println(makeGood("leEeetcode"))
	fmt.Println(makeGood("abBAcC"))
	fmt.Println(makeGood("s"))
}