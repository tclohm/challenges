package main

import "fmt"

func sortJumbled(mapping, nums []int) []int {
	// ht := map[int]int{}
	for _, n := range nums {
		mappedN := 0
		base := 1
		for n > 0 {
			digit := n % 10
			n = n / 10
			mappedN += base * mapping[digit]
			base *= 10
			fmt.Println(digit, n, mappedN, base)
		}
	}
	return []int{}
}

func main() {
	fmt.Println(sortJumbled([]int{8,9,4,0,2,1,3,5,7,6}, []int{991,338,38}))
	fmt.Println(sortJumbled([]int{0,1,2,3,4,5,6,7,8,9}, []int{789,456,123}))
}
