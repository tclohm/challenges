package main

import "fmt"

func concat(nums []int) []int {
	res := []int{}
	for i := 0 ; i < 2 ; i++ {
		for _, n := range nums {
			res = append(res, n)
		}
	}
	return res
}

func main() {
	fmt.Println(concat([]int{1,2,1}))
}