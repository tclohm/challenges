package main

import (
	"fmt"
)

func findLonely(nums []int) []int {
	var result = []int{}
	var ht = map[int]int{}

	for i := 0 ; i < len(nums) ; i++ {
		n := nums[i]
		if _, exist := ht[n]; !exist {
			ht[n] = 1
		} else {
			ht[n] += 1
		}
	}

	for _, n := range nums {
		_, ok := ht[n-1]
		_, exist := ht[n+1]
		if ht[n] == 1 && ok == false && exist == false {
			result = append(result, n)
		}
	}

	return result
}

func main() {
	fmt.Println(findLonely([]int{10,6,5,8}))
	fmt.Println(findLonely([]int{1,3,3,5}))
}