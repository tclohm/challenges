package main

import (
	"fmt"
	"strings"
)

type stack struct {
	container []string
}

func minRemoveToMakeValid(s string) string {
	result := []string{}
	count := 0

	for _, r := range s {
		if r == '(' {
			result = append(result, string(r))
			count++
		} else if r == ')' && count > 0 {
			result = append(result, string(r))
			count--
		} else if r != ')' {
			result = append(result, string(r))
		}
	}

	filtered := []string{}
	for i := len(result) - 1 ; i >= 0 ; i-- {
		if result[i] == "(" && count > 0 {
			count -= 1
		} else {
			filtered = append(filtered, result[i])
		}
	}

	for i, j := 0, len(filtered) - 1 ; i < j ; i, j = i+1, j-1 {
		filtered[i], filtered[j] = filtered[j], filtered[i]
	}

	return strings.Join(filtered, "")
}

func main() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
	fmt.Println(minRemoveToMakeValid("a)b(c)d"))
	fmt.Println(minRemoveToMakeValid("))(("))
}