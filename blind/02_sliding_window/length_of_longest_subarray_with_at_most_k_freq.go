package main

import "fmt"

// [1,2,3,1,2,3,1,2] k = 3
// x = number of time it occurs in an array
// array is good if the freq of each element is less than or equal to k
// return the length of the longest good subarray

func maxSubarrayLength(nums []int, k int) int {
	var length int = 0
	var count = map[string]int{}
	var left int = 0
	
	for right, _ := range nums {
		char := string(nums[right])
		count[char]++
		for count[char] > k {
			count[string(nums[left])] -= 1
			left++
		}
		length = max(length, right - left + 1)
	}

    return length
}

func main() {
	fmt.Println(maxSubarrayLength([]int{1,2,3,1,2,3,1,2}, 2))
	fmt.Println(maxSubarrayLength([]int{1,2,1,2,1,2,1,2}, 1))
	fmt.Println(maxSubarrayLength([]int{5,5,5,5,5,5,5}, 4))
}