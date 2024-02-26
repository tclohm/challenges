package main 

import "fmt"

// nums[i] = max lenght you can jump forward
// 0 <= j <= nums[i]
// return the min number of jumps to reach nums[n - 1]

func jump(nums []int) int {
	//var n = len(nums)
	var min = 1000
	return min
}

func minimum(a, b int) int {
	if a < b { return a }
	return b
}

func main() {
	fmt.Println(jump([]int{2,3,1,1,4}))
	fmt.Println(jump([]int{3,2,0,1,4}))
}