package main

import "fmt"

func hasDuplicates(nums []int) bool {
	ht := map[int]int{}
	for _, n := range nums {
		if ht[n] > 0 { return true }
		ht[n]++
	}
	return false
}

func main() {
	fmt.Println(hasDuplicates([]int{1,2,3,3}))
fmt.Println(hasDuplicates([]int{1,2,3,4}))
}
