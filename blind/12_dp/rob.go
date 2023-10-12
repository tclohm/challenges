package main

import "fmt"

func rob(nums []int) int {
	// 2 options: rob current house or don't rob current house
	// rob(i) = max(rob(i - 2) + currentValue, rob(i - 1))
	return robbing(nums, len(nums) - 1)
}

func robbing(nums []int , index int) int {
	if index < 0 {
		return 0
	}
	return max(robbing(nums, index - 2) + nums[index], robbing(nums, index - 1))
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func robbingMemo(nums []int) int {
	memo := make([]int, len(nums) + 1)

	for i := 0 ; i < len(memo) ; i++ {
		memo[i] = -1
	}

	return robMemo(nums, len(nums) - 1, memo)
}

func robMemo(nums []int, index int, memo []int) int {
	if index < 0 {
		return 0
	}
	if memo[index] >= 0 {
		return memo[index]
	}
	// index - 2, index - 1 is always comparing against each other
	// index - 2 always adds nums[index]
	// recursive call on both using the stack to keep stored results
	result := max(robMemo(nums, index - 2, memo) + nums[index], robMemo(nums, index - 1, memo))
	memo[index] = result
	return result
}

func bottomUpRob(nums []int) int {
	if len(nums) == 0 { return 0 }
	memo := make([]int, len(nums) + 1)
	memo[0] = 0
	memo[1] = nums[0]

	for i := 1 ; i < len(nums) ; i++ {
		value := nums[i]
		memo[i+1] = max(memo[i], memo[i-1] + value)
	}
	return memo[len(nums)]
}

func main() {
	fmt.Println(bottomUpRob([]int{1,2,3,1}))
	fmt.Println(bottomUpRob([]int{2,7,9,3,1}))
}