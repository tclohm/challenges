package main

import "fmt"

// word or phrase rearrange the letters of a different word or phrase

func validHTS(s string, t string) bool {
	if len(s) != len(t) { return false }
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


func valid(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var alpha [26]int

	for i := 0 ; i < len(s) ; i++ {
		alpha[s[i] - 'a']++
		alpha[t[i] - 'a']--
	}

	for i := 0 ; i < len(alpha) ; i++ {
		if alpha[i] != 0 { return false }
	}

	return true
}

func main() {
	fmt.Println(valid("anagram", "nagaram"))
	fmt.Println(valid("rat", "car"))
}