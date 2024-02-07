package main

import "fmt"


// compare finger prints
func checkInclusion(s1, s2 string) bool {
	left, count := 0, [26]int{}
	for i := range s1 { count[s1[i]-'a']++ }

	for right := range s2 {
		count[s2[right]-'a']--
		if count == [26]int{} { return true }

		if right + 1 >= len(s1) {
			count[s2[left]-'a']++
			left++
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}