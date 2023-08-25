package main

import "fmt"

// word or phrase rearrange the letters of a different word or phrase

func valid(s string, t string) bool {
	ht := make(map[rune]int, len(s))

	collect(ht, s)
	remove(ht, t)

    for _, v := range ht {
    	if v > 0 { return false }
    }

    return true
}

func collect(ht map[rune]int, s string) {
	for _, r := range s {
		if _, ok := ht[r]; ok {
			ht[r]++
		} else {
			ht[r] = 1
		}
	}
}

func remove(ht map[rune]int, t string) {
	for _, r := range t {
		if _, ok := ht[r]; ok {
			ht[r]--
		} else {
			ht[r] = 1
		}
	}
}

func main() {
	fmt.Println(valid("anagram", "nagaram"))
	fmt.Println(valid("rat", "car"))
}