package main

import "fmt"

func intersection(nums1, nums2 []int) []int {
	ht1 := map[int]bool{}
	ht2 := map[int]bool{}

	result := []int{}

	for _, d := range nums1 {
		ht1[d] = true
	}

	for _, d := range nums2 {
		ht2[d] = true
	}

	for k, _ := range ht1 {
		if _, ok := ht2[k]; ok {
			result = append(result, k)
		}
	}

	return result
}

func main() {
	fmt.Println(intersection([]int{4,9,5}, []int{9,4,9,8,4}))
	fmt.Println(intersection([]int{1,2,2,1}, []int{2,2}))
}