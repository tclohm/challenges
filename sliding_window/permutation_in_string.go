package main

import "fmt"

func check(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	s1Count, s2Count := make([]rune, 26, 26), make([]rune, 26, 26)
	
	for i, r := range s1 {
		s1Count[r - 'a'] += 1
		s2Count[rune(s2[i]) - 'a'] += 1
	}

	//fmt.Println(s1Count, s2Count)
	matches := 0

	for i := 0 ; i < 26 ; i++ {
		if s1Count[i] == s2Count[i] {
			matches += 1
		}
	}

	leftPointer := 0 
	for rightPointer := len(s1) ; rightPointer < len(s2) ; rightPointer++ {
		if matches == 26 {
			return true
		}

		var index = rune(s2[rightPointer]) - 'a'
		s2Count[index] += 1

		if s1Count[index] == s2Count[index] {
			matches += 1
		} else if s1Count[index] + 1 == s2Count[index] {
			matches -= 1
		}

		index = rune(s2[leftPointer]) - 'a'
		s2Count[index] -= 1

		if s1Count[index] == s2Count[index] {
			matches += 1
		} else if s1Count[index] - 1 == s2Count[index] {
			matches -= 1
		}

		leftPointer += 1
	}
	
	return matches == 26
}

func main() {
	fmt.Println(check("ab", "eidbaooo"))
	fmt.Println(check("ab", "eidboaoo"))
}