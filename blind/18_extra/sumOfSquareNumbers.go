package main

import (
	"fmt"
	"math"
)

func bf(result int) bool {
	for i := 1 ; i < result ; i++ {
		for j := 1 ; j < result ; j++ {
			if i * i + j * j == result {
				return true
			}
 		}
	}
	return false
}

func isSumOfSquareNumbers(result int) bool {
	squareRoot := map[int]bool{}
	for i := 1 ; i < int(math.Floor(math.Sqrt(float64(result)))) ; i++ {
		squareRoot[i * i] = true
	}

	a := 0
	for a * a <= result {
		target := result - a * a
		if _, ok := squareRoot[target]; ok {
			return true
		}
		a++
	}
	return false
}

func main() {
	fmt.Println(isSumOfSquareNumbers(5))
	fmt.Println(isSumOfSquareNumbers(3))
}
