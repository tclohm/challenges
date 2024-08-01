package main

import "fmt"

func subarraySum(nums []int, k int) int {
	total := 0
	left, right := make([]int, len(nums)), make([]int, len(nums))	
	
	left[0] = nums[0]
	for i := 1 ; i < len(nums) ; i++ {
		left[i] = left[i - 1] + nums[i] 
	}
	
	end := len(nums) - 1
	right[end] = nums[end]

	for i := end - 1 ; i >= 0 ; i-- {
		right[i] = right[i + 1] + nums[i]
	}
	
	left = append(left, right...)

	for i := range left {
		if left[i] == k {
			total++
		}
	}

	return total
}

func sub(nums []int, k int) int {
	var sum int
	var count int
	var prevSum = make(map[int]int)
	prevSum[0] = 1
	for i := range nums {
		sum += nums[i]		
		if val, ok := prevSum[sum - k]; ok {
			count += val
		}
		prevSum[sum]++
	}
	return count
}

func main() {
	fmt.Println(subarraySum([]int{1,1,1}, 2))
	fmt.Println(subarraySum([]int{1,2,3}, 3))

}
