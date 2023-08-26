package main

import "fmt"

func twoSum(nums []int, target int) []int {
    ht := make(map[int]int, len(nums))

    for index, value := range nums {
    	remainder := target - value
    	if position, ok := ht[remainder]; ok {
    		return []int{position, index}
    	} 
    	ht[value] = index
    }

    return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{3,2,4}, 6))
	fmt.Println(twoSum([]int{3,3}, 6))
}