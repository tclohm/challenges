package main

import (
	"fmt"
	"math"
)

func leftRightDifference(nums []int) []int {
	prefix := make([]int, len(nums), len(nums))
	prefix[0] = nums[0]
	for i := 1 ; i < len(nums) ; i++ {
		prefix[i] = prefix[i - 1] + nums[i]
	}
	
	postfix := make([]int, len(nums), len(nums))
	end := len(nums) - 1
	postfix[end] = nums[end]

	for i := end - 1 ; i >= 0 ; i-- {
		postfix[i] = postfix[i + 1] + nums[i]
	}

	for i := 0 ; i < len(nums) ; i++ {
		diff := postfix[i] - prefix[i]
		prefix[i] = int(math.Abs(float64(diff)))
	}

	return prefix
}

func lrd(nums []int) []int {
	res := make([]int, len(nums))
	l, r := 0, 0
	for _, n := range nums {
		r += n
	}
	for i, n := range nums {
		r -= n
		res[i] = int(math.Abs(float64(l - r)))
		l += n
	}
	return res
}

func main() {
	fmt.Println(lrd([]int{10,4,8,3}))
	fmt.Println(lrd([]int{1}))
}
