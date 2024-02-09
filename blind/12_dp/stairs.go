package main

import "fmt"

func climb(steps int) int {
	memo := map[int]int{}

	var decision func(steps int) int
	decision = func(steps int) int {
		if steps == 1 {
			return 1
		}
		if steps == 2 {
			return 2
		}

		memo[steps - 1] = decision(steps - 1)
		memo[steps - 2] = decision(steps - 2)

		return memo[steps - 1] + memo[steps - 2]
	}

	return decision(steps)
}

func main() {
	fmt.Println(climb(2))
	fmt.Println(climb(3))
	fmt.Println(climb(4))
	fmt.Println(climb(5))
}