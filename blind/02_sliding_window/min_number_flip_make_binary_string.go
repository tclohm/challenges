package main

import "fmt"

func min(a,b,c int) int {
	if a < b && a < c {
		return a
	}
	if b < a && b < c {
		return b
	}
	return c
}

func minFlips(s string) int {
	n := len(s)
	s = s + s

	alt1, alt2 := "", ""

	for i, _ := range s {
		if i % 2 == 0 {
			alt1 += "0"
			alt2 += "1"
		} else {
			alt1 += "1"
			alt2 += "0"
		}
	}

	res := len(s)
	diff1, diff2 := 0, 0
	l := 0

	for r, _ := range s {
		if s[r] != alt1[r] {
			diff1 += 1
		}

		if s[r] != alt2[r] {
			diff2 += 1
		}

		if (r - l + 1) > n {
			if s[l] != alt1[l] {
				diff1 -= 1
			}
			if s[l] != alt2[l] {
				diff2 -= 1
			}
			l += 1
		}

		if (r - l + 1) == n {
			res = min(res, diff1, diff2)
		}
	}

	return res
}

func main() {
	fmt.Println(minFlips("111000"))
	fmt.Println(minFlips("010"))
	fmt.Println(minFlips("1110"))
}