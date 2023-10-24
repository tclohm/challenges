package main

import "fmt"

func missing(nums []int) int {
	var res int = len(nums)

	for i := 0 ; i < len(nums) ; i++ {
		res += i - nums[i]
	}

	return res
}

func main() {
	fmt.Println(missing([]int{3,0,1}))
	fmt.Println(missing([]int{0,1}))
	fmt.Println(missing([]int{9,6,4,2,3,5,7,0,1}))
}