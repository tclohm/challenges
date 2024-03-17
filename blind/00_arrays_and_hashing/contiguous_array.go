package main

import "fmt"

func findMaxLength(nums []int) int {
	zero, one := 0, 0

	diff_index := map[int]int{}

	result := 0

    for i, d := range nums {
    	if d == 0 {
    		zero++
    	} else {
    		one++
    	}

    	if _, ok := diff_index[one - zero]; !ok {
    		diff_index[one - zero] = i
    	}

    	if one == zero {
    		result = one + zero
    	} else {
    		idx := diff_index[one - zero]
    		result = max(result, i - idx)
    	}
    }
    return result
}

func main() {
	fmt.Println(findMaxLength([]int{0,1}))
	fmt.Println(findMaxLength([]int{0,1,0}))
}