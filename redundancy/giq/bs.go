package main

import "fmt"

func bs(arr []int, target int) int {
	left, right := 0, len(arr) - 1

	for left <= right {

		middle := left + (right - left) / 2

		fmt.Println(arr[middle], middle)

		if arr[middle] > target {
			right = middle - 1
		} else if arr[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}

	return -1
}

func main() {
	sorted := []int{12, 14, 16, 18, 22, 43, 66, 83, 94, 105}
	fmt.Println(bs(sorted, 66))
}