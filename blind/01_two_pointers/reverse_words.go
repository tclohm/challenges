package main

import (
	"fmt"
	"strings"
)

func reverse(s string) string {
	words := strings.Fields(s)

	left, right := 0, len(words) - 1

	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	return strings.TrimSpace(strings.Join(words, " "))
}

func rev(s string) string {
	result := ""

	for l := 0 ; l < len(s) ; l++ {
		if s[l] == ' ' { continue }

		r := l
		// we found a word
		for r < len(s) && s[r] != ' ' {
			r++
		}

		if result == "" {
			result = s[l:r]
		} else {
			result = s[l:r] + " " + result
		}
		l = r
	}

	return result
}

func main() {
	fmt.Println(rev("the sky is blue"))
	fmt.Println(rev("  hello world  "))
	fmt.Println(rev("a good   example"))
}