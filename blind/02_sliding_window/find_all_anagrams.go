package main

import "fmt"

func mapsEqual(ht1, ht2 map[rune]int) bool {
	for k1, v1 := range ht1 {
		if v2, ok := ht2[k1]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

func findAnagrams(haystack, target string) []int {
	// find the starting indices dor the anagrams
	target_ht := make(map[rune]int)
	haystack_ht := make(map[rune]int)

	result := []int{}

	for _, r := range target {

		target_ht[r]++

	}

	left := 0

	for right, right_char := range haystack {

		haystack_ht[right_char]++

		window := right - left + 1

		for window == len(target) {

			left_char := rune(haystack[left])

			if mapsEqual(target_ht, haystack_ht) {

				result = append(result, left)

			}

			if haystack_ht[left_char] > 1 {

				haystack_ht[left_char]--

			} else {

				delete(haystack_ht, left_char)

			}

			left++
			window = right - left + 1

		}
	}

	return result
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}