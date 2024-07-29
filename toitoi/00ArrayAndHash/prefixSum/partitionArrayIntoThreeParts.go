package main

import (
	"fmt"
	// "math"
)

func sum(nums []int) int {
	total := 0
	for _, n := range nums { 
		total += n
	}
	return total
}

func avg(total int) int {
	return total / 3
}

func canThreePartsEqualSum(nums []int) bool {
	sum := sum(nums)
	if sum % 3 != 0 { return false }
	l, r, avg := 0, len(nums) - 1, avg(sum)
	fmt.Println(avg)
	leftSum, rightSum := nums[l], nums[r]
	for l <= r {
		if l < r && leftSum != avg {
			leftSum += nums[l]
			l += 1
		}
		if l < r && rightSum != avg {
			rightSum += nums[r]
			r -= 1
		}
		if leftSum == avg && avg == rightSum {
			return true
		}
		if l == r {
			return false
		}
	}
	return false
}

func main() {
	fmt.Println(canThreePartsEqualSum([]int{0,2,1,-6,6,-7,9,1,2,0,1}))
	fmt.Println(canThreePartsEqualSum([]int{0,2,1,-6,6,7,9,-1,2,0,1}))
	fmt.Println(canThreePartsEqualSum([]int{3,3,6,5,-2,2,5,1,-9,4}))
}
