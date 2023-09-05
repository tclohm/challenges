package main

import "fmt"

type stack struct {
	container []string
}

func isValid(s string) bool {
	st := stack{}
	for _, r := range s {
		str := string(r)
		st.container = append(st.container, str)
	}
	
	return false
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
}