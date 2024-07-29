package main

import "fmt"

func sum(nums []int) int {
	total := 0
	for _, n := range nums { 
		total += n
	}
	return total
}

func avg(total, n int) int {
	return total / n
}

func canThreePartsEqualSum(nums []int) bool {
	sum := sum(nums)
	N := len(nums)
	if sum / 3 != 0 { return false }
	l, r, avg := 0, len(nums) - 1, avg(sum, N)
	fmt.Println(l, r, avg)
	return true
}

func main() {
	fmt.Println(canThreePartsEqualSum([]int{0,2,1,-6,6,-7,9,1,2,0,1}))
	fmt.Println(canThreePartsEqualSum([]int{0,2,1,-6,6,7,9,-1,2,0,1}))
	fmt.Println(canThreePartsEqualSum([]int{3,3,6,5,-2,2,5,1,-9,4}))
}
