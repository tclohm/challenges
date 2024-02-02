package main

import "fmt"

func minWindow(str string, target string) string {

	if len(target) > len(str) || target == "" { return "" }

	target_ht := map[string]int{}
	window := map[string]int{}

	for i := 0 ; i < len(target) ; i++ {
		if _, ok := target_ht[string(target[i])] ; !ok {
			target_ht[string(target[i])] = 1
		} else {
			target_ht[string(target[i])] += 1
		}
	}

	have, need := 0, len(target_ht)
	result, result_length := []int{-1, -1}, 10000
	l := 0
	// start with right pointer and collect the letters
	for r, _ := range str {
		char := string(str[r])
		if _, ok := window[char]; ok {
			window[char] += 1
		} else {
			window[char] = 1
		}

		if _, ok := target_ht[char]; ok && window[char] == target_ht[char] {
			have += 1
		}

		for have == need {

			if r - l + 1 < result_length {
				result = []int{l, r}
				result_length = r - l + 1
			}

			left_char := string(str[l])


			window[left_char] -= 1

			if _, ok := target_ht[left_char]; ok && window[left_char] < target_ht[left_char] {
				have -= 1
			}

			l += 1
		}
	}

	return str[result[0]:result[1]+1]
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("a", "aa"))
}