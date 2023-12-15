package main

import "fmt"
// O(n) time and space
func shuffle(nums []int, k int) []int {
	result := []int{}
	for i := 0 ; i < k ; i++ {
		result = append(result, nums[i])
		result = append(result, nums[i + k])
	}
	return result
}
// O(1) time and space
func shuffleInPlace(nums []int, k int) []int {
	// will return to this
}

func main() {
	fmt.Println(shuffleInPlace([]int{2,5,1,3,4,7}, 3))
}