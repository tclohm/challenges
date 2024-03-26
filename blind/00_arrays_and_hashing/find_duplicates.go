package main

import (
	"fmt"
	"math"
)

func findDuplicates(nums []int) []int {
	var results = []int{}

	for _, n := range nums {
		absolute := math.Abs(float64(n))

		if nums[int(absolute) - 1] < 0 {
			results = append(results, int(absolute))
		}

		nums[int(absolute) - 1] = nums[int(absolute )- 1] * -1
	}
	return results
}

func main() {
	fmt.Println(findDuplicates([]int{4,3,2,7,8,2,3,1}))
}