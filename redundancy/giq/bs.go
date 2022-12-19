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

func bsWithPointers(arr []int, left, right, target int) int {
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

func startAndEnd(arr []int, target int) []int {
	// binary to find target initially
	// use left and right pointers to find end
	if len(arr) == 0 { return []int{-1, -1} }
	midPoint := bs(arr, target)
	var startPoint = midPoint
	var endPoint = midPoint
	var tmp int
	for startPoint != -1 {
		tmp = startPoint
		startPoint = bsWithPointers(arr, 0, startPoint - 1, target)
	}

	startPoint = tmp

	for endPoint != -1 {
		tmp = endPoint
		endPoint = bsWithPointers(arr, endPoint + 1, len(arr) - 1, target)
	}

	endPoint = tmp

	return []int{startPoint, endPoint}

}

func main() {
	sorted := []int{12, 14, 16, 18, 22, 43, 66, 83, 94, 105}
	fmt.Println(bs(sorted, 66))

	sorted2 := []int{12, 14, 16, 17, 17, 17, 18, 22, 43, 66, 83, 94, 105, 105, 105, 105, 105, 105, 105, 105}
	fmt.Println(startAndEnd(sorted2, 106))
}