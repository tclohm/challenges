package main

import "fmt"

func sub(nums []int, k int) int {
	var (
		prefixSum = 0
		result = 0
		prefixCount = make(map[int]int)
	)
	prefixCount[0] = 1

	for n := range nums {
		prefixSum += n
		remain := prefixSum % k

		if _, ok := prefixCount[remain]; ok {
			result += prefixCount[remain]
		}
		prefixCount[remain]++
	}
	return result
}

func main() {
	fmt.Println(sub([]int{4,5,0,-2,-3,1}, 5))
	fmt.Println(sub([]int{5}, 9))

}
