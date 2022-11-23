package main

import (
	"fmt"
	"sort"
)



type set map[[3]int]bool

func threeSum(n []int) [][]int {
	// sort the integers
	var result = [][]int{}

	if len(n) == 0 || len(n) == 1 {
		return result
	}

	sort.Ints(n[:])

	for index, integer := range n {
		if index > 0 && integer == n[index - 1] {
			continue
		}
		// grab two pointers and add them together with the integer in iteration
		// if the sum is above zero right pointer decrements
		// if the sum is less than zero left pointer increments
		// else add the numbers to the result array in its own array
		left, right := index + 1, len(n) - 1

		for left < right {
			three_sum := integer + n[left] + n[right]

			if three_sum > 0 {
				right -= 1
			} else if three_sum < 0 {
				left += 1
			} else {
				result = append(result, []int{integer, n[left], n[right]})
				left += 1
				
				for n[left] == n[left - 1] && left < right {
					left += 1
				}
			}
		}
	}

	return result

}

func main() {
	////////////////////////
	// -4, -1, -1, 0, 1, 2
	///////////////////////
	// integer : -4 , left : -1, right : 2  && three_sum = -3
	// integer : -4,  left : -1, right : 2  && three_sum = -3
	// integer : -4,  left : 0,  right : 2  && three_sum = -2
	// integer : -4,  left : 1,  right : 2  && three_sum = -1
	// integer : -4,  left : 2,  right : 2  LEFT AND RIGHT ARE SAME INDEX

	// MOVE TO NEXT INTEGER
	// integer : -1,  left : 1,  right : 1  && three_sum = -2

	n := []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(n))
}