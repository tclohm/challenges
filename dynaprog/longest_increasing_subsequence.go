package main

import "fmt"

// integer array nums, return the length of the longest strictly increasing subsequence
// (
// subsequence is an array that can be derived 
// from another array by deleting some or no 
// elements without changing the order of the remaining
// )

// how do we compute the longest possible sequence at a given index? 
// The way to do that is to look at maximum lengths of sequences at all previous 
// indices and find the highest one that corresponds to a value that is lower then the current one

// ex: [10, 9, 2, 5, 3, 7, 101, 18]
// lisLength[0] = 1
// lisLength[1] = 1
// lisLength[2] = 1
// lisLength[3] = ?

// iteratively, look at the previous values in lisLengths
//	a. if the corresponding value is lower (index 2, input[2] < input[3]) then that's a candidate
//     for finding the longest previous sequence
//  b. check if we haven't found a higher one earlier and if we haven't, set the value

// overlapping subproblems, dp problem
func lengthOfLIS(nums []int) int {
	lisLengths := make([]int, len(nums))

	for i := range lisLengths {
		lisLengths[i] = 1
	}

	for i := 1 ; i < len(nums) ; i++ {
		for j := 0 ; j < i ; j++ {
			if nums[j] < nums[i] && lisLengths[j] + 1 > lisLengths[i] {
				lisLengths[i] = lisLengths[j] + 1
			}
		}
	}

	var max int
	for _, length := range lisLengths {
		if length > max {
			max = length
		}
	}

	return max
}

func main() {
	fmt.Println(lengthOfLIS([]int{10,9,2,5,3,7,101,18}))
	fmt.Println(lengthOfLIS([]int{0,1,0,3,2,3}))
	fmt.Println(lengthOfLIS([]int{7,7,7,7,7,7,7}))
}