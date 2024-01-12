package main

import (
	"math"
	"fmt"
)

// two types of operations
// two elements with equal values can be deleted from the array
// three elements with equal values and deleted from the array

// 1 item in the array --- -1
// 2 item in the array --- 1 operation
// 3 item in the array --- 1 operation

func Constructor(nums []int) map[int]int {
	ht := make(map[int]int)

	for _, n := range nums {
		if _, ok := ht[n]; ok {
			ht[n] += 1
		} else {
			ht[n] = 1
		}
	}

	return ht
}

func minOperation(nums []int) int {
	count := Constructor(nums)
	fmt.Println("ht", count)
	var operations float64 = 0

	for _, c := range count {
		if c == 1 {
			return -1
		}
		// 3 is largest number
		operations += math.Ceil(float64(c / 3))
	}
	return int(operations)
}


func main() {

	nums := []int{2,3,3,2,2,4,2,3,4}
	fmt.Println(minOperation(nums))
}