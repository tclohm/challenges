package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1
   
   	for start != end {
   		if numbers[end] + numbers[start] == target { return []int{start+1, end+1} }
   		if numbers[end] + numbers[start] < target { start++ }
   		if numbers[end] + numbers[start] > target { end-- }
   	}
    	
    return []int{0,0}
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{2,3,4}, 6))
	fmt.Println(twoSum([]int{-1, 0}, -1))
}