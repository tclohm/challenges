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
func lengthOfLIS(nums []int) int {

}

func main() {
	fmt.Println(lengthOfLIS([]int{10,9,2,5,3,7,101,18}))
	fmt.Println(lengthOfLIS([]int{0,1,0,3,2,3}))
	fmt.Println(lengthOfLIS([]int{7,7,7,7,7,7,7}))
}