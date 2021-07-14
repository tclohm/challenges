package main
import (
	//"math/rand"
	"fmt"
)

func quicksort(array []int, start int, end int) {
	if start >= end {
		return
	}
	index := partition(array, start, end)
	fmt.Println(index)
	quicksort(array, start, index-1)
	quicksort(array, index+1, end)
}

func partition(array []int, start int, end int) int {
	pivotIndex := start
	pivotValue := array[end]

	for i := start; i < end; i++ {
		if array[i] < pivotValue {
			swap(array, i, pivotIndex)
			pivotIndex++
		}
	}

	swap(array, pivotIndex, end)
	return pivotIndex
}

func swap(array []int, a int, b int) {
	temp := array[a]
	array[a] = array[b]
	array[b] = temp
}

func main() {
	fixedArray := []int{10, 23, 4, 5, 6, 2, 9, 11, 3, 1, 14, 32}

	fmt.Println(fixedArray)
	quicksort(fixedArray, 0, len(fixedArray) - 1)
	fmt.Println("quick sorted!")
	fmt.Println(fixedArray)
}