package main

import "fmt"

func appendCharacters(s string, t string) int {
	var count = 0
	var i = 0
	for j, _ := range t {
		if t[j] == s[i] && j == i {
			i++
		} else {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(appendCharacters("coaching", "coding"))
	fmt.Println(appendCharacters("abcde", "a"))
	fmt.Println(appendCharacters("z", "abcde"))
}
