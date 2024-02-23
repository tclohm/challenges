package main

import "fmt"

// n people
// 1 - n
// judge : trusts no one
// everyone trust the judge
func findJudge(n int, trust [][]int) int {
	judges := make(map[int]int, n+1)
	for _, array := range trust {
		self := array[0]
		other := array[1]
		judges[self]--
		judges[other]++
	}

	for i, judge := range judges {
		if judge == n - 1 { return i }
	}
	return -1
}

func main() {
	fmt.Println(findJudge(2, [][]int{{1,2}})) // 2
	fmt.Println(findJudge(3, [][]int{{1,3},{2,3}})) // 3
	fmt.Println(findJudge(3, [][]int{{1,3},{2,3},{3,1}})) // -1
}