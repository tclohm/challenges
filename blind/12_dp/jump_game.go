package main 

import "fmt"

// greedy
func canJump(nums []int) bool {
	goalpost := len(nums) - 1
	for i := len(nums) - 1 ; i >= 0 ; i-- {
		if i + nums[i] >= goalpost {
			goalpost = i
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

func min(a, b int) int {
	if a < b { return a }
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

func can_jump(nums []int) bool{
	var can_jump_from_position func(position_index int) bool
	can_jump_from_position = func(position_index int) bool {
		if position_index == len(nums) - 1 {
			return true
		}
		// furthest jump from current position
		furthest_jump := min(position_index + nums[position_index], len(nums) - 1)
		// try all possible jumps from the current position
		for next_position_index := position_index + 1 ; next_position_index < furthest_jump ; next_position_index++ {
			if can_jump_from_position(next_position_index) {
				return true
			}
		}
		return false
	}
	return can_jump_from_position(0)
}

func main() {
	fmt.Println(dp_can_jump([]int{2,3,1,1,4}))
	fmt.Println(dp_can_jump([]int{3,2,1,0,4}))
}

// 3 => 1: 2, 2: 1, 3: 0
//		   2 => 1: 1, 2: 0
//				1 => 1: 0

// Greed algos make local optimal choices at each stage 
// with the hope of finding global optimum