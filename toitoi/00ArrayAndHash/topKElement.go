package main

import "fmt"

func topK(nums []int, k int) []int {
	// the k most
	count := make(map[int]int)
	frequency := make([][]int, len(nums) + 1, len(nums) + 1)
	
	top := []int{}
	for _, n := range nums {
		count[n]++
	}
	// save the number to the frequencies index in the array
	for number, c := range count {
		frequency[c] = append(frequency[c], number)
	}
	// start from the back of our frequency array and grab the biggest frequencies
	for i := len(frequency) - 1 ; i > 0 ; i-- {
		for _, n := range frequency[i] {
			top = append(top, n)
			if len(top) == k { return top }
		}
	}


	return top
}

func main() {
	fmt.Println(topK([]int{1,2,2,3,3,3,5,5,5}, 2))
	fmt.Println(topK([]int{7,7}, 1))
}
