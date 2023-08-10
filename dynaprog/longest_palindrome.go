package main

import "fmt"

func longest(str string) string {
	result := ""
	length := 0

	for i, _ := range str {
		// odd length
		left, right := i, i
		for left >= 0 && right < len(str) && str[left] == str[right] {
			if (right - left + 1) > length {
				result = str[left:right + 1]
				length = right - left + 1
			}
			left -= 1
			right += 1
		}

		// even length
		left, right = i, i + 1
		for left >= 0 && right < len(str) && str[left] == str[right] {
			if (right - left + 1) > length {
				result = str[left:right + 1]
				length = right - left + 1
			}
			left -= 1
			right += 1
		}
	}
	return result
}

func main() {
	fmt.Println(longest("babad"))
	fmt.Println(longest("racecar"))
	fmt.Println(longest("cbbd"))
}