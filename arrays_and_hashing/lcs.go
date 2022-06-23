package main

import "fmt"

func lcs(nums []int) int {
	hm := make(map[int]bool)

	for _, i := range nums {
		hm[i] = true
	}

	fmt.Println(hm)

	return 0	 
}

func main() {
	n := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(lcs(n))
}