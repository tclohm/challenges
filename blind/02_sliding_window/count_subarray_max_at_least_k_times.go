package main

import "fmt"

func max(ns []int) int {
	big, number := 0, 0
	ht := make(map[int]int)
	for _, n := range ns {
		ht[n]++
		if big < ht[n] {
			number = n
		}
	}
	return number
}

func countSubarrays(nums []int, k int) int {
   max_n, max_freq := max(nums), 0
   left, count := 0, 0
   for right, _ := range nums {
   	if nums[right] == max_n {
   		max_freq++
   	}
   	for max_freq == k {
   		if nums[left] == max_n {
   			max_freq--
   		}
   		left++
   	}
   	count++
   }
   return count
}

func main() {
	fmt.Println(countSubarrays([]int{1,3,2,3,3}, 2))
	fmt.Println(countSubarrays([]int{1,4,2,1}, 3))
}