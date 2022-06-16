package main

import (
	"fmt"
)

// Replace Elements with Greatest Element on Right Side
// array, replace every element in that array with the greatest element
// among the elements to its right, and replace the last element with -1

// ex arr = [17, 18, 5, 4, 6, 1]
// output = [18, 6, 6, 6, -1]

// new[0] = max(arr[1:5])
// new[1] = max(arr[2:5])
// ...

// new[0] = max(arr[i], new[1])
// reverse the order

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// initial maximum = -1
	// reverse iteration
	// new max = max(oldmax, arr[i])
	var arr = []int{17, 18, 5, 4, 6, 1}
	var rightMax int = -1

	for i := len(arr) - 1 ; i > -1 ; i-- {
		// find new max
		var newMax = max(rightMax, arr[i])
		// replace i with the right max
		arr[i] = rightMax
		// update right max
		rightMax = newMax
	}

	fmt.Println(arr)
}