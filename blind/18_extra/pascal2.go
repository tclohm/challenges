package main

import "fmt"

func getRow(rowIndex int) []int {
	array := make([]int, rowIndex + 1)
	for i := range array {
		array[i] = 1
	}

	var previous, current int

	for i := 2 ; i <= rowIndex ; i++ {
		previous = 1
		for k := 1 ; k < i ; k++ {
			current = array[k]
			array[k] = array[k] + previous
			previous = current
		}
	}

	return array
}

func main() {
	fmt.Println(getRow(1))
	fmt.Println(getRow(2))
	fmt.Println(getRow(3))
	fmt.Println(getRow(4))
	fmt.Println(getRow(5))
}