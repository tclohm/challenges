package main

import "fmt"

// given array of integers nums and integer, k
// continous subarray is called nice if there are
// k odd numbers on it

// sliding window variation
// - outer loop increments window right pointer until `k` odd numbers are found
// - `k` odd numbers are found, in inner loop, increment our nice count
// - add the number of nice subarrays to result
func numberOfSubarrays(nums []int, k int) int {
	count := 0
	left, middle := 0, 0
	for right, _ := range nums {
		if nums[right] % 2 != 0 {
			odd++
		}

		for odd == k {
			if nums[left] % 2 != 0 {
				odd--
			}
			left++
			middle = left
		}

		if odd == k {
			for nums[middle] % 2 != 0 {
				middle++
			}
			count += (m - l) + 1
		}
	}
	return count
}

func main() {
	fmt.Println(numberOfSubarray([]int{1,1,2,1,1}, 3))
	fmt.Println(numberOfSubarray([]int{2,4,6}, 1))
}
