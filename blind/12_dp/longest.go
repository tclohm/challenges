package main

import "fmt"

func longest(nums []int) int {
	cache := make([]int, len(nums))
	LIS := 0
	for i := range nums {
		currentMax := 0
		for j := 0 ; j < i + 1 ; j++ {
			if nums[j] < nums[i] {
				if currentMax < cache[j] { currentMax = cache[j] }
			}
		}
		cache[i] = currentMax + 1
		if cache[i] > LIS {
			LIS = cache[i]
		}
	}
	return LIS
}

func main() {
	fmt.Println(longest([]int{10,9,2,5,3,7,101,18}))
	fmt.Println(longest([]int{0,1,0,3,2,3}))
	fmt.Println(longest([]int{7,7,7,7,7,7,7}))
}