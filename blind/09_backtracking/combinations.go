package main

import "fmt"

// k numbers chosen
// from range [1,n]
func combine(n, k int) [][]int {
	var list [][]int

	var backtrack func(combination []int, index int)
	backtrack = func(combination []int, index int) {
		if len(combination) == k {
			list = append(list, append([]int{}, combination...))
			return
		}

		for i := index ; i <= n ; i++ {
			combination = append(combination, i)
			backtrack(combination, i+1)
			combination = combination[:len(combination) - 1]
		}
	}

	backtrack([]int{}, 1)
	return list
}

func main() {
	fmt.Println(combine(4,2))
	fmt.Println(combine(1,1))
}