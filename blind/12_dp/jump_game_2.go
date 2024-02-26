package main 

import "fmt"

// nums[i] = max length you can jump forward
// 0 <= j <= nums[i]
// return the min number of jumps to reach nums[n - 1]

func jump(nums []int) int {
	var n = len(nums) - 1
	var dfs func(position int) int
	dfs = func(position int) int {
		var min = 10001
		if position >= n { return 0 }
		// why the 1? because the
		for i := 1 ; i <= nums[position] ; i++ {
			min = minimum(min, 1 + dfs(i + position))
		}
		return min
	}
	return dfs(0)
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func minimum(a, b int) int {
	if a < b { return a }
	return b
}

func jumpGreedy(nums []int) int {
	min := 0
	left, right := 0, 0

	for right < len(nums) - 1 {
		farthest := 0
		for i := left ; i < right + 1 ; i++ {
			farthest = max(farthest, i + nums[i])
		}
		left = right + 1
		right = farthest
		min += 1
	}
	return min
}

func main() {
	fmt.Println(jumpGreedy([]int{2,3,1,1,4}))
	fmt.Println(jumpGreedy([]int{3,2,0,1,4}))
	fmt.Println(jumpGreedy([]int{1,1,1,1,1,1,1,1}))
	fmt.Println(jumpGreedy([]int{1,2,2,1,1,2,2,1}))
}

// 2 => [3,1]
// 		3 => [1,1,4]
// 		   1 => [1]
// 		  		1 => [4]
//
// 		   1 => [1]
//				 1 => [4]
