package main

import "fmt"

func countBits(n int) []int {
	count := make([]int, n+1)
	for i := 1 ; i <= n ; i++ {
		if i & 1 == 1 {

		}
		count[i] = count[i >> 1] + i
	}
	return count
}

func main() {
	fmt.Println(countBits(2))
}