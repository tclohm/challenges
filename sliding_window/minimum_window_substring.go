package main

import "fmt"

func minWindow(s, t string) string {
	if t == "" {
		return ""
	}

	countT, window := map[rune]int{}, map[rune]int{}

	for _, r := range t {
		if _, exist := countT[r]; exist {
			countT[r] += 1
		} else {
			countT[r] = 1
		}
	}

	have, need := 0, len(countT)
	result := []int{-1, -1}
	resultLength := -100000
	lp := 0

	for rp := 0 ; rp < len(s) ; rp++ {
		c := s[rp]
		if _, exist := window[c]; exist {
			window[c] += 1
		} else {
			window[c] = 1
		}

		_, ok := countT[c]

		if window[c] == countT[c] && ok {
			have += 1
		}

		for have == need {
			if (rp - lp + 1) < resultLength {
				result = {lp, rp}
				resultLength = (rp - lp + 1)
			}
			window[s[l]] -= 1
		}
	}
}

func main() {

}