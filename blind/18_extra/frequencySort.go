package main

import "fmt"

func frequency(nums []int) []int {
	freq := map[int]int{}
	for _, n := range nums {
		freq[n]++
	}


	return []int{} 
}


func main() {
	fmt.Println(frequency([]int{1,1,2,2,2,3}))
	fmt.Println(frequency([]int{2,3,1,3,2}))
	fmt.Println(frequency([]int{-1,1,-6,4,5,-6,1,4,1}))
}
