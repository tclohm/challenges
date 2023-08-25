package main

import "fmt"

func contains(nums []int) bool {
	ht := make(map[int]bool, len(nums))
	
	for _, n := range nums {
		if _, ok := ht[n]; ok {
			return ht[n]
		} else {
			ht[n] = true
		}
	}
	return false
}

func main() {
	fmt.Println(contains([]int{1, 2, 3, 1}))
	fmt.Println(contains([]int{1, 2, 3, 4}))
	fmt.Println(contains([]int{1,1,1,3,3,4,3,2,4,2}))
}