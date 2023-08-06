package main

import (
	"fmt"
)

// climb stairs, each time you can take 1 or 2 steps. How many distinct ways can you climb to the top?
func climb(stairs int) int {
	var cache = make([]int, stairs+1)
	return climbing(stairs, cache)
}

func climbing(stairs int, cache []int) int {
	if stairs < 3 {
		return stairs
	}

	if cache[stairs] != 0 {
		return cache[stairs]
	}

	result := climbing(stairs - 1, cache) + climbing(stairs - 2, cache)
	cache[stairs] = result
	return result
}

func main() {
	fmt.Println(climb(2))
	fmt.Println(climb(45))
}