package main

import "fmt"

func selectionSort(elements []int) {
	for i := 0 ; i < len(elements) - 1 ; i++ {
		var min int = i

		for j := i + 1 ; j <= len(elements) - 1 ; j++ {
			if elements[j] < elements[min] {
				min = j
			}
		}
		swap(elements, i, min)
	}
}

func swap(elements []int, i int, j int) {
	var tmp int = elements[j]
	elements[j] = elements[i]
	elements[i] = tmp
}

func main() {
	var elements []int
	elements = []int{11, 4, 18, 6, 19, 21, 71, 13, 15, 2}
	fmt.Println("Before sorting ", elements)
	selectionSort(elements)
	fmt.Println("After sorting", elements)
}