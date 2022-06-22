package main

import "fmt"

func productExceptSelf(numbers []int) []int {
	var res = []int{}

	for i := 0 ; i < len(numbers) ; i++ {
		var product int = 1 
		for j := 0 ; j < len(numbers) ; j++ {
			if i != j {
				product *= numbers[j]
			}
		}
		res = append(res, product)
	}

	return res
}

func pes(n []int) []int {
	var res = make([]int, len(n), len(n))
	var prefix = 1
	for i := 0 ; i < len(n) ; i++ {
		res[i] = prefix
		prefix *= n[i]
	}
	var postfix int = 1
	for i := len(n) - 1 ; i > -1 ; i-- {
		res[i] *= postfix
		postfix *= n[i]
	}

	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	nums2 := []int{-1, 1, 0, -3, 3}

	fmt.Println(pes(nums))
	fmt.Println(pes(nums2))
}