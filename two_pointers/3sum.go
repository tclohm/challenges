package main

import (
	"fmt"
	"sort"
)



type set map[[3]int]bool

func threeSum(n []int) [][]int {
	var result = [][]int{}
	//var s = set{}
	if len(n) == 0 || len(n) == 1 {
		return result
	}
	sort.Ints(n[:])


	for i, a := range n {
		if i > 0 && a == n[i - 1] {
			continue
		}

		l, r := i + 1, len(n) - 1

		for l < r {
			tsum := a + n[l] + n[r]

			if tsum > 0 {
				r -= 1
			} else if tsum < 0 {
				l += 1
			} else {
				result = append(result, []int{a, n[l], n[r]})
				l += 1
				for n[l] == n[l - 1] && l < r {
					l += 1
				}
			}
		}
	}

	return result

}

func main() {
	// -4, -1, -1, 0, 1, 2
	n := []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(n))
}