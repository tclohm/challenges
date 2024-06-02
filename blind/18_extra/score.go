package main

import "fmt"

func score(s string) int {
	result := 0
	prev := s[0]
	for i := 1 ; i < len(s) ; i++ {
		diff := abs(int(prev) - int(s[i]))
		result += diff
		prev = s[i]
	}
	return result
}

func abs(a int) int {
	if a < 0 { return -a }
	return a
}

func main() {
	fmt.Println(score("hello"))
	fmt.Println(score("zaz"))
}

