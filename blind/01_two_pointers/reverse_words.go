package main

import (
	"fmt"
	"strings"
)

func reverse(words string) string {
	left, right := 0, len(words) - 1

	arr := strings.Split(words, "")
	// reverse the positions
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}

	

	return strings.Join(arr, "")
}

func main() {
	fmt.Println(reverse("the sky is blue"))
	fmt.Println(reverse("  hello world  "))
	fmt.Println(reverse("a good   example"))
}