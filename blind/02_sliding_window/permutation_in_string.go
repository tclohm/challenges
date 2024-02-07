package main

import "fmt"


// compare finger prints
func checkInclusion(target, haystack string) bool {
	want := [26]int{}
	for _, r := range target { 
		want[r -'a']++ 
	}

	window := len(target)
	have := [26]int{}
	// build up an array, if the arrays are similar, we have a match
	for right, char := range haystack {
		// count tells us how many of the current chars we have seen in the window
		have[char-'a']++
		// decrement chars out of the window
		if right >= window {
			outOfWindow := haystack[right - window]
			have[outOfWindow - 'a']--
		}
		// if our arrays are equal return true
		if have == want { return true }
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}