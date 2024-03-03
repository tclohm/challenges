package main

import (
	"fmt"
	"strings"
)

func maxOddBinaryNumber(s string) string {
	arr := make([]string, len(s), len(s))
	left := 0

	for i, r := range s {
		arr[i] = string(r)
	}

	for i, r := range s {
		if string(r) == "1" {
			arr[i], arr[left] = arr[left], arr[i]
			left += 1
		}
	}
	arr[left - 1], arr[len(arr) - 1] = arr[len(arr) - 1], arr[left - 1]
	return strings.Join(arr, "")
}

func main() {
	fmt.Println(maxOddBinaryNumber("010"))
	fmt.Println(maxOddBinaryNumber("1001"))
}