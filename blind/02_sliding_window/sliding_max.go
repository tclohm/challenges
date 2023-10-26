package main

import "fmt"


func maxSlidingWindow(nums []int, k int) []int {
    result := []int{}

    for i := 0 ; i + k <= len(nums) ; i++ {
    	max := max(nums[i:i+k])
    	result = append(result, max)
    }
    return result
}

func max(n []int) int {
	if len(n) == 1 { return n[0] }
	big := n[0]
	for i := 0 ; i < len(n) ; i++ {
		if big < n[i] {
			big = n[i]
		}
	}
	return big
}


func main() {
	fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3))
	fmt.Println(maxSlidingWindow([]int{1}, 1))
}