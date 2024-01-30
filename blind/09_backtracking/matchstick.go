package main

import (
	"fmt"
	"sort"
)

func sum(nums []int) int {
	total := 0 
	for _, n := range nums {
		total += n
	}
	return total
}

func makesquare(matchsticks []int) bool {
	length := sum(matchsticks) / 4
	sides := make([]int, 4, 4)

	if sum(matchsticks) / 4 != length {
		return false
	}

	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))

	var backtrack func (index int) bool
	backtrack = func (index int) bool {
		if index == len(matchsticks) {
			return true
		}

		for jndex := 0 ; jndex < 4 ; jndex++  {
			if sides[jndex] + matchsticks[index] <= length {
				sides[jndex] += matchsticks[index]
				if backtrack(index + 1) {
					return true
				}
				sides[jndex] -= matchsticks[index]
			}
		}
		return false
	}

	return backtrack(0) 
}

func main() {
	fmt.Println(makesquare([]int{1,1,2,2,2}))
	fmt.Println(makesquare([]int{3,3,3,3,4}))
}