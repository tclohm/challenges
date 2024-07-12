package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {

	var stack = []string{}

	for _, r := range s {
		c := string(r)
		if c == ")" {
			portion := []string{}
			for stack[len(stack) - 1] != "(" {

				last := len(stack) - 1
				el := stack[last]
				stack = stack[:last]

				portion = append(portion, el)
			}
			last := len(stack) - 1
			stack = stack[:last]
			stack = append(stack, portion...)
		} else {
			stack = append(stack, c)
		}
	}
	return strings.Join(stack, "")
}

func main() {
	fmt.Println(reverse("(abcd)"))
	fmt.Println(reverse("(u(love)i)"))
  fmt.Println(reverse("(ed(et(oc))el)"))
}

func rev(s string) string {
	pair := map[int]int
	stack := []string{}

	for i, r := range s {
		c := string(r)
		if c == "(" {
			stack = append(stack, c)
		} else if c == ")" {
			last := len(stack) - 1
			el := stack[last]
			stack = stack[:last]
			pair[i] = el
			pair[el] = j
		}
	}

	index, direction := 0, 1
	result := []string{}

	for index < len(s) {
		if s[index] == "(" || s[index] == ")" {
			index = pair[index]
			direction = -direction
		} else {
			result = append(result, s[i])
		}
		i += direction
	}

	return strings.Join(result, "")
}
