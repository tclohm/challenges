package main

import "fmt"

func sum(nums []int) int {
	total := 0
	for _, n := range nums { 
		total += n
	}
	return total
}

func canThreePartsEqualSum(arr []int) bool {
	total := sum(arr)
	fmt.Println(total)
	return true
}

func main() {
	fmt.Println(canThreePartsEqualSum([]int{0,2,1,-6,6,-7,9,1,2,0,1}))
	fmt.Println(canThreePartsEqualSum([]int{0,2,1,-6,6,7,9,-1,2,0,1}))
	fmt.Println(canThreePartsEqualSum([]int{3,3,6,5,-2,2,5,1,-9,4}))
}
