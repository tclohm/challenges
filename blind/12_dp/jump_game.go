package main 

import "fmt"

// greedy
func canJump(nums []int) bool {
	goal := len(nums) - 1
	for i := len(nums) - 1 ; i >= 0 ; i-- {
		if i + nums[i] >= goal {
			goal = i
		}
	}

	return goal == 0
}

// idea of this solution
// get the max index we can go up to by getting the max of i + n[i] and max
// the if i has surpassed the max index, we can return false
// otherwise return true


// updates the last_position to the current index if you can reach 
// the last position from the current index. Finally, it checks if last_position is 0, 
// meaning you can reach the last index from the first index.
func jump(n []int) bool {
	var i = 0
	for reach := 0 ; i < len(n) && i <= reach ; i++ {
		reach = max(i + n[i], reach)
	}

	return i == len(n)
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func dp_can_jump(nums []int) bool {
	n := len(nums)

	dp := make([]bool, n, n)

	dp[n - 1] = true

	for i := n - 2 ; i > -1 ; i-- {
		// is it possible to reach the last index from the current index
		for j := 1 ; j < nums[i] + 1 ; j++ {
			if i + j < n && dp[i + j] == true {
				dp[i] = true
				break
			}
		}
	}

	return dp[0]
}

func main() {
	fmt.Println(dp_can_jump([]int{2,3,1,1,4}))
	fmt.Println(dp_can_jump([]int{3,2,1,0,4}))
}

// Greed algos make local optimal choices at each stage 
// with the hope of finding global optimum