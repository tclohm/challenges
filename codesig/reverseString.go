package main
import (
	"fmt"
	"strings"
)

func reverseString(args string) string {
	words := strings.Split(args, " ")
	reversed := ""
	for i := len(words) - 1 ; i >= 0 ; i-- {
		reversed += " " + words[i]
	}

	return reversed[1:]
}

func main() {
	words := "hello world"
	fmt.Println(words)
	reversed := reverseString(words)
	fmt.Println(reversed)

	newwords := "Why do I have fear over things that are out of my control"
	fmt.Println(newwords)
	reversenNewWords := reverseString(newwords)
	fmt.Println(reversenNewWords)
}