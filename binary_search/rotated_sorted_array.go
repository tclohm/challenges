package main

import "fmt"

func bruteForceSearch(nums []int, target int) int {
	for index, number := range nums {
		if number == target {
			return index
		}
	}
	return -1
}


func main() {
	n1 := []int{4,5,6,7,0,1,2}

	fmt.Println(bruteForceSearch(n1, 0))

	fmt.Println(bruteForceSearch(n1, 3))

	n2 := []int{1}

	fmt.Println(bruteForceSearch(n2, 0))
}