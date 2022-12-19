package main

import "fmt"

func qs(arr *[]int, left, right int) []int {
	if left <= right {
		partitionPoint := partition(arr, left, right)
		qs(arr, left, partitionPoint - 1)
		qs(arr, partitionPoint + 1, right)
	}

	return *arr
}

func partition(arr *[]int, left, right int) int {
	var pivotElement = (*arr)[right]
	var partitionPoint = left

	for i := left ; i < right ; i++ {
		if (*arr)[i] < pivotElement {
			// swap
			(*arr)[i], (*arr)[partitionPoint] = (*arr)[partitionPoint], (*arr)[i]
			partitionPoint++
		}
	}

	(*arr)[right], (*arr)[partitionPoint] = (*arr)[partitionPoint], (*arr)[right]
	return partitionPoint
}

func getKthElement(arr []int, k int) int {
	index := len(arr) - k
	if index > len(arr) || index < 0 {
		return -1
	}
	qs(&arr, 0, len(arr) - 1)
	return arr[index]
}


func main() {
	unsorted := []int{2, 3, 1, 2, 4, 2}
	un1 := []int{5, 3, 1, 6, 4, 2}
	un2 := []int{4, 2, 3, 2, 4, 0}
	un3 := []int{3}

	// fmt.Println(qs(&unsorted, 0, len(unsorted) - 1))
	// fmt.Println(qs(&un1, 0, len(un1) - 1))
	// fmt.Println(qs(&un2, 0, len(un2) - 1))
	// fmt.Println(qs(&un3, 0, len(un3) - 1))

	fmt.Println(getKthElement(unsorted, 1))
	fmt.Println(getKthElement(un1, 5))
	fmt.Println(getKthElement(un2, 1))
	fmt.Println(getKthElement(un3, 1))
}