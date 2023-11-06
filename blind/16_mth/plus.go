package main

import "fmt"

func plus(digits []int) []int {
	for i := len(digits) - 1 ; i >= 0 ; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		}
	}
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}

func main() {
	fmt.Println(plus([]int{1,2,3}))
	fmt.Println(plus([]int{4,3,2,1}))
	fmt.Println(plus([]int{9}))
}
