package main

import (
	"fmt"
	"math"
)
func matches(n float64) int {
	var res float64 = 0
	for n > 1 {
		res += n / 2
		n = math.Ceil(n / 2)
	}
	return int(res)
}

func matchesE(n int) int {
	return n - 1
}
func main() {
	fmt.Println(matchesE(7))
}