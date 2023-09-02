package main

import "fmt"

func length_of_longest_substring(s string) int {
	ht := make(map[byte]int)
	longest := 0
	length := 0
	for i := 0 ; i < len(s) ; i++ {
		b := s[i]
		if _, exist := ht[b]; exist {
			if length >= longest {
				longest = length
				length = 0
			}
			ht[b] = i
			length += 1
		} else {
			ht[b] = i
			length += 1
		}
	}
	return longest
}

func main() {
	fmt.Println(length_of_longest_substring("abcabcbb"))
	fmt.Println(length_of_longest_substring("bbbbb"))
	fmt.Println(length_of_longest_substring("pwwkew"))
}