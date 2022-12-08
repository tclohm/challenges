package main

import "fmt"

func typedOut(s, t string) bool {
	var result, result2 = "", ""
	var i = len(s) - 1

	for i >= 0 {
		if s[i] == '#' {
			i--
		} else {
			result += string(s[i])
		}
		i--
	}

	i = len(t) - 1

	for i >= 0 {
		if t[i] == '#' {
			i--
		} else {
			result2 += string(t[i])
		}
		i--
	}

	return result == result2
}

func main() {
	a1 := "ab#z"
	a2 := "az#z"

	b1 := "az"
	b2 := "az"

	c1 := "ab#z"
	c2 := "az#z"

	fmt.Println(typedOut(a1, a2))
	fmt.Println(typedOut(b1, b2))
	fmt.Println(typedOut(c1, c2))
}