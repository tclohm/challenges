package main

import "fmt"

func maxLength(chars []string) int {
	var set = map[rune]int{}

	var overlap func (set map[rune]int, s string) bool
	overlap = func (set map[rune]int, s string) bool {
		var ht = map[rune]int{}
		for _, r := range s {
			_, ok := set[r]
			_, exist := ht[r]
			if ok || exist {
				return true
			}
			ht[r] = 1
		}
		return false
	}	

	var backtrack func (index int) int
	backtrack = func (index int) int {
		if index == len(chars) {
			return len(set)
		}
		var result = 0
		if !overlap(set, chars[index]) {
			for _, letter := range chars[index] {
				if _, ok := set[letter]; !ok {
					set[letter] = 1
				} else {
					set[letter] += 1
				}
				
			}
			result = backtrack(index + 1)
			for _, letter := range chars[index] {
				delete(set, letter)
			}
		}
		return max(result, backtrack(index + 1))
	}

	return backtrack(0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func main() {
	fmt.Println(maxLength([]string{"un", "iq", "ue"}))
	fmt.Println(maxLength([]string{"cha", "r", "act", "ers"}))
}