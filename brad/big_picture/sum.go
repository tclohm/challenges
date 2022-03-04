package main

import (
	"fmt"
	"time"
)

// Linear O(n)
func sum(to int) int{
	start := time.Now()
	total := 0
	for i := 0 ; i < to + 1 ; i++ {
		total += i
	}
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("time elapsed", elapsed)
	
	return total
}

// O(1)
func arithmetic_sum(to int) int {
	start := time.Now()
	total := to * (to + 1) / 2
	end := time.Now()
	fmt.Println("time elapsed", end.Sub(start))
	return total
}

func main() {
	fmt.Println("sum(4) ->", sum(4), "\n")
	fmt.Println("sum(10000) ->", sum(10000), "\n")
	fmt.Println("sum(100000000) ->", sum(100000000), "\n")

	fmt.Println("faster sum(4)", arithmetic_sum(4), "\n")
	fmt.Println("faster sum(10000)", arithmetic_sum(10000), "\n")
	fmt.Println("faster sum(100000000)", arithmetic_sum(100000000), "\n")
}