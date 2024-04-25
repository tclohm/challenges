package main

import (
	"math"
	"fmt"
)

func longestIdealString(s string, k int) int {
	var result int
	var dfs func(index int, previous string) int
	dfs = func(index int, previous string) int {
		if index == len(s) { return 0 }


		// skip
		result = dfs(index + 1, previous)

		
		if previous != "" {
			index_number := s[index] - 'a'
			prev_number := previous - 'a'
			if int(math.Abs(float64(index_number) - float64(prev_number))) <= k {
				result = max(result, dfs(index + 1, string(s[index])))
			}
		}

		return result
	}

	return dfs(0, "")
}


func main() {
	fmt.Println(longestIdealString("acfgbd", 2))
	fmt.Println(longestIdealString("abcd", 3))
}