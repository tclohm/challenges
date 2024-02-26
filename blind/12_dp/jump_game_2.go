package main 

import "fmt"

// nums[i] = max lenght you can jump forward
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

func main() {
	fmt.Println(jump([]int{2,3,1,1,4}))
	fmt.Println(jump([]int{3,2,0,1,4}))
	fmt.Println(jump([]int{1,1,1,1,1,1,1,1}))
	fmt.Println(jump([]int{1,2,2,1,1,2,2,1}))
}

// 2 => [3,1]
// 		3 => [1,1,4]
// 		   1 => [1]
// 		  		1 => [4]
//
// 		   1 => [1]
//				 1 => [4]
