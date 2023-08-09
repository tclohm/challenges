package main

import (
	"fmt"
)

// climb stairs, each time you can take 1 or 2 steps. How many distinct ways can you climb to the top?
// Time: O(n)
// Space: O(n)
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

// Bottom up
// Time: O(n)
// Space: O(n)

func climbAgain(n int) int {
	if n < 3 {
		return n
	}

	ways := make([]int, n+1)
	ways[1] = 1
	ways[2] = 2

	for i := 3 ; i <= n ; i++ {
		ways[i] = ways[i-1] + ways[i-2]
	}

	return ways[n]
}

// Bottom up (Optimized)
// Time: O(n)
// Space: O(1)

func climbOptimized(n int) int {
	if n < 3 {
		return n
	}

	prev, current := 1, 2
	for i := 3 ; i <= n ; i++ {
		prev, current = current, prev+current
	}

	return current
}

func main() {
	fmt.Println(climbOptimized(2))
	fmt.Println(climbOptimized(45))
}