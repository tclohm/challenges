package main

import "fmt"

func intersection(nums1, nums2 []int) []int {
	seen := map[int]bool{}
	result := []int{}

	for _, d := range nums1 {
		seen[d] = true
	}

	for _, d := range nums2 {
		if seen[d] == true {
			result = append(result, d)
			seen[d] = false
		}
	}

	return result
}

func main() {
	fmt.Println(intersection([]int{4,9,5}, []int{9,4,9,8,4}))
	fmt.Println(intersection([]int{1,2,2,1}, []int{2,2}))
}