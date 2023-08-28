package main

import (
	"fmt" 
	"strings"
)

// Design an algorithm to encode a list of strings to a string. 
// The encoded string is then sent over the network and is decoded 
// back to the original list of strings.
func encode(strs []string) string {
	result := ""

	for i, w := range strs {
		if i == len(strs) - 1 {
			result += w
		} else {
			result += w + ":;"
		}
	}

	return result
}

func decode(str string) []string {
	return strings.Split(str, ":;")
}

func main() {
	enc1 := encode([]string{"lint","code","love","you"})
	enc2 := encode([]string{"we", "say", ":", "yes"})
	fmt.Println("encoded", enc1, enc2)
	fmt.Println(decode(enc1))
	fmt.Println(decode(enc2))
}