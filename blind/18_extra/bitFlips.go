package main

import "fmt"

func minK(nums []int, k int) int {
	var (
		queue = []int{}
		res = 0
	)

	for i, _ := range nums {

		for len(queue) > 0 && queue[0] + k - 1 {
			queue = queue[1:]
		}


		if nums[i] + len(queue) % 2 == 0 {
			if i + k > len(nums) { 
				return -1 
			}
			res += 1
			queue = append(queue, i)
		}
	}

	return res
}

func main() {
	fmt.Println(minK([]int{0,1,0}, 1))
	fmt.Println(minK([]int{1,1,0}, 2))
}
