package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
    total := 0
    left := 0
    product := 1
    for right, value := range nums {
    	product *= value
    	for left <= right && product >= k {
    		product = product / nums[left]
    		left += 1
    	}
    	total += (right - left + 1)
    }
    return total
}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10,5,2,6}, 100))
	fmt.Println(numSubarrayProductLessThanK([]int{1,2,3}, 0))
}