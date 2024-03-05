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

func oddBinary(s string) string {
	count := 0 
	for _, r := range s {
		if r == '1' {
			count += 1
		}
	}

	str := ""
	for i := 0 ; i < count - 1 ; i++ {
		str += "1"
	}

	for i := 0 ; i < len(s) - count ; i++  {
		str += "0"
	}

	return str + "1"
}

func main() {
	fmt.Println(oddBinary("010"))
	fmt.Println(oddBinary("1001"))
}