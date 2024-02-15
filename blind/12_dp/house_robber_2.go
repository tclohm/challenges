package main

import "fmt"

func robDP(nums []int) int {
	if len(nums) == 0 { return 0 }
	if len(nums) == 1 { return nums[0] }

    var robbing func(n []int) int
    robbing = func(n []int) int {
    	if len(n) == 0 { return 0 }
    	if len(n) == 1 { return nums[0] }

    	dp := make([]int, len(n) + 1)
    	dp[0] = 0
    	dp[1] = n[0]

    	for i := 1 ; i < len(n) ; i++ {
    		dp[i+1] = max(dp[i], n[i] + dp[i - 1])
    	}

    	return dp[len(n)]
    }

    return max(robbing(nums[1:]), robbing(nums[:len(nums) - 1]))
}

func rob(nums []int) int {
	if len(nums) == 0 { return 0 }
	if len(nums) == 1 { return nums[0] }

	var robbing func(nums []int) int
	robbing = func(nums []int) int {
		prev := 0
		curr := nums[0]

		for i := 1 ; i < len(nums) ; i++ {
			tmp := prev
			prev = max(prev, curr)
			curr = tmp + nums[i]
		}
		return max(prev, curr)
	}
	

	return max(robbing(nums[1:]), robbing(nums[:len(nums) - 1]))
}

func max(a, b int) int {
    if a > b { return a }
    return b
}


func main() {
	fmt.Println(rob([]int{2,3,2}))
	fmt.Println(rob([]int{1,2,3,1}))
}