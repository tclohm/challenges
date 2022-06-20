package main

// o(log n) -- most likely binary solution

import "fmt"

func search(rotated_array []int, target int) int {
	var (
		left = 0
		right = len(rotated_array) - 1
	)

	for left <= right {
		middle := (left + right) / 2

		if target == rotated_array[middle] {
			return middle
		}

		if rotated_array[left] <= rotated_array[middle] {

			if target > rotated_array[middle] || target < rotated_array[left] {
				left = middle + 1
			} else {
				right = middle - 1
			}

		} else {

			if target < rotated_array[middle] || target > rotated_array[right] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return -1
}

func main() {
	numbers := []int{4, 5, 6, 7, 0, 1, 2,}
	target := 0
	fmt.Println(search(numbers, target))
}