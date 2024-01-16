package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	trimmed := strings.TrimSpace(s)
	array := strings.Split(trimmed, " ")
	return len(array[len(array) - 1])
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}