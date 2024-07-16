package main

import (
	"fmt"
	"strings"
)

func maxGain(s string, x, y int) int {
	var removePairs func (pair string, score int) int
	removePairs = func (pair string, score int) int {
		var result = 0
		var stack = []string{}

		for _, r := range s {
			var c = string(r)
			if c == string(pair[1]) && len(stack) > 0 && stack[len(stack) - 1] == string(pair[0]) {
				stack = stack[:len(stack) - 1]
				result += score
			} else {
				stack = append(stack, c)
			}
		}
		return result
	}

	var res = 0
	var pair string
	if x > y {
		pair = "ab"
	}	else {
		pair = "ba"
	}
	
	var rev = reverse(pair)

	res += removePairs(pair, max(x, y))
	res += removePairs(rev, min(x, y))
	return res
}

func reverse(s string) string {
	var res strings.Builder
	for i := len(s) - 1 ; i >= 0 ; i-- {
		res.WriteString(string(s[i]))
	}
	return res.String()
}

func max(a, b int) int {
	if a < b { return b }
	return b
}

func min(a, b int) int {
	if a < b { return a }
	return b
}

func main() {
	fmt.Println(maxGain("cdbcbbaaabab", 4, 5)) // 19
	fmt.Println(maxGain("aabbaaxybbaabb", 5, 4)) // 20
}
