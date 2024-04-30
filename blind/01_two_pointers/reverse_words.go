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

func main() {
	fmt.Println(reverse("the sky is blue"))
	fmt.Println(reverse("  hello world  "))
	fmt.Println(reverse("a good   example"))
}