package main

import "fmt"

func isSumOfSquareNumbers(result int) bool {
	for i := 1 ; i < result ; i++ {
		for j := 1 ; j < result ; j++ {
			if i * i + j * j == result {
				return true
			}
 		}
	}
	return false
}

func main() {
	fmt.Println(isSumOfSquareNumbers(5))
	fmt.Println(isSumOfSquareNumbers(3))
}
