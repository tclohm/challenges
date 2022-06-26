package main

import "fmt"

func twoSum(nums []int, target int) []int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		sum := nums[start] + nums[end]
		
		if sum == target {
			return []int{start+1, end+1}
		}

		if sum < target {
			start += 1
		} else {
			end -= 1
		}
	}

	return []int{0, 0}
}

func main() {
	n1 := []int{2,7,11,15}
	n2 := []int{2,3,4}
	n3 := []int{-1,0}

	fmt.Println(twoSum(n1, 9))
	fmt.Println(twoSum(n2, 6))
	fmt.Println(twoSum(n3, -1))
}