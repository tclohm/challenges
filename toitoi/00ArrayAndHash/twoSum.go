package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// [number] = index 
	ht := map[int]int{}
	for index, num := range nums {
		if _, exist := ht[target - num]; exist {
			return []int{ht[target - num], index}
		}
		ht[num] = index
	}
	return []int{-1,-1}
}

func main() {
	fmt.Println(twoSum([]int{3,4,5,6}, 7))
	fmt.Println(twoSum([]int{4,5,6}, 10))
	fmt.Println(twoSum([]int{5,5}, 10))
}
