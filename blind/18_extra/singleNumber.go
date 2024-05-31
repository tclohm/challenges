package main

import "fmt"

func single(nums []int) []int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}

	setBit := xor & -xor
	a, b := 0, 0
	for _, n := range nums {
		if n & setBit != 0 {
			a ^= n
		} else {
			b ^= n
		}
	}
	return []int{a,b}
}

func main() {
	fmt.Println(single([]int{1,2,1,3,2,5}))
	fmt.Println(single([]int{-1, 0}))
	fmt.Println(single([]int{0,1}))
}
