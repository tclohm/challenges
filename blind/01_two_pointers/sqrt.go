package main

import (
	"fmt"
	"math"
)

func sqroot(x int) int {
	left, right := 0, x
	result := 0

	for left <= right {
		mid := left + ((right - left) / 2)

		if math.Pow(float64(mid), 2) > float64(x) {
			right = mid - 1
		} else if math.Pow(float64(mid), 2) < float64(x) {
			left = mid + 1
			result = mid
		} else {
			return mid
		}
	}
	return result
}

func main() {
	fmt.Println(sqroot(4))
	fmt.Println(sqroot(50))
	fmt.Println(sqroot(100))
}