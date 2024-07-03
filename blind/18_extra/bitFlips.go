package main

import "fmt"

func minK(nums []int, k int) int {
	var (
		currentWindowFlips = 0
		res = 0
	)

	for i, _ := range nums {
		if i - k >= 0 && nums[i - k] == 2 {
			currentWindowFlips -= 1
		}

		if nums[i] + currentWindowFlips % 2 == 0 {
			if i + k > len(nums) { 
				return -1
			}
			res += 1
			currentWindowFlips += 1
			nums[i] = 2
		}
	}
	return res
}

func main() {
	fmt.Println(minK([]int{0,1,0}, 1))
	fmt.Println(minK([]int{1,1,0}, 2))
}
