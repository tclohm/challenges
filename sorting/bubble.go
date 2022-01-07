package main

import "fmt"

func bubbleSort(integers [11]int) {
	num := 11
	isSwapped := true

	for isSwapped {
		isSwapped = false
		for i := 1 ; i < num ; i++ {
			if integers[i - 1] > integers[i] {
				tmp := integers[i]
				integers[i] = integers[i - 1]
				integers[i - 1] = tmp
				isSwapped = true
			}
		}
	}
	fmt.Println(integers)
}

func main() {
	var integers [11]int = [11]int{31, 13, 12, 4, 18, 16, 7, 2, 3, 0, 10}
	fmt.Println("Sort")
	bubbleSort(integers)
}