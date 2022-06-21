package main

import "fmt"

// {
// 	"a": 0
// 	"n": 0
// 	"g": 0
// 	"r": 0
// 	"m": 0
// }

func isAnagram(s, t string) bool {

	if len(s) != len(t) {
		return false
	}

	hm := make(map[byte]int)
	for i := 0 ; i < len(s) ; i++ {
		b := s[i]
		_, ok := hm[b]
		if ok {
			hm[b] += 1
		} else {
			hm[b] = 1
		}
	}

	for i := 0 ; i < len(t) ; i++ {
		b := t[i]
		v, ok := hm[b]
		if ok && v > 0 {
			hm[b] -= 1
		} else {
			return false
		}
	}
	
	for _, value := range hm {
		if value != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "cat"))
	fmt.Println(isAnagram("racecar", "carrace"))
}