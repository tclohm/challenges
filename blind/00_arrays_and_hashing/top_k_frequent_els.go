package main

import ("fmt"
		"math")


func topKFrequent(nums []int, k int) []int {
	max := biggest(nums)
	result := []int{}

    frequency := make([]int, max+1)
    for _, n := range nums {
    	frequency[n]++
    }

    for index, count := range frequency {
    	if count >= k {
    		result = append(result, index)
    	}
    }
   
    return result
}

func biggest(nums []int) int {
	var biggest int = int(math.Inf(-1))
	for _, number := range nums {
		if number > biggest {
			biggest = number
		}
	}

	return biggest
}


func main() {
	fmt.Println(topKFrequent([]int{1,1,1,2,2,3}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
	fmt.Println(topKFrequent([]int{1,2,2,3,3,3,4,4,5,5,5,6,6,6,7,7,8,8,8,9,9,9,9}, 3))
}