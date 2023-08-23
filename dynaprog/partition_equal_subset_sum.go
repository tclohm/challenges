package main

import "fmt"


// given an array, return true if you can partition the array into two subsets
// sum of the elements in bot subsets is equal or false otherwise
func canPartition(nums []int) bool {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	if sum % 2 == 1 { return false }

	sum = sum / 2

	dp := make([]bool, sum, sum) // reachable numbers
	dp[0] = true

	for _, n := range nums {
		if n <= sum { // n must be less than or equal to sum
			if dp[sum - n] == true { // found sum
				return true
			}

			// loop in opposite direction, because 
			// we don't want to check index and 
			// then loop over it
			for k := sum - n - 1 ; k >= 0 ; k-- {
				if dp[k] == true {
					dp[k + n] = true
				}
			}
		}
	}

	return false   
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
	fmt.Println(canPartition([]int{1, 2, 4, 3}))

}