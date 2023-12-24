package main

import "fmt"

func removeElement(nums []int, val int) int {
	var length int = 0
	for _, n := range nums {
		if n == val { continue }
		length++
	}
	return length
}

func main() {
	fmt.Println(removeElement([]int{3,2,2,3}, 3))
	fmt.Println(removeElement([]int{0,1,2,2,3,0,4,2}, 2))
}