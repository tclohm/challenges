package main

import "fmt"

func checkInclusion(s1, s2 string) bool {
	if len(s1) > len(s2) { return false }

	s1Count, s2Count := make([]rune, 26, 26), make([]rune, 26, 26)

	for i, _ := range s1 {
		s1Count[s1[i] - 'a']++
		s2Count[s2[i] - 'a']++
	}

	matches := 0
	for i := 0 ; i < 26 ; i++ {
		if s1Count[i] == s2Count[i] {
			matches++
		}
	}

	left := 0
	for right := len(s1) ; right < len(s2) ; right++ {
		if matches == 26 { return true }

		index := rune(s2[right]) - 'a'
		s2Count[index]++

		if s1Count[index] == s2Count[index] {
			matches += 1
		} else if s1Count[index] + 1 == s2Count[index] {
			matches -= 1
		}

		index = rune(s2[left]) - 'a'
		s2Count[index]--

		if s1Count[index] == s2Count[index] {
			matches += 1
		} else if s1Count[index] - 1 == s2Count[index] {
			matches -= 1
		}

		left++
	}

	return matches == 26
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}