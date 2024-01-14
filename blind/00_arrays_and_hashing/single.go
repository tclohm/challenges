package main

import "fmt"

func single(nums []int) int {
	ht := make(map[int]int)
	for _, n := range nums {
		if _, ok := ht[n]; ok {
			ht[n] += 1
		} else {
			ht[n] = 1
		}
	}

	for key, value := range ht {
		if value == 1 {
			return key
		}
	}
	return -1
}

func main() {
	fmt.Println(single([]int{2,2,1}))
	fmt.Println(single([]int{4,1,2,1,2}))
	fmt.Println(single([]int{1}))
}