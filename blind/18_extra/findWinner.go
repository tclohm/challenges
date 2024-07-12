package main

import "fmt"

func findWinner(n, k int) int {
	q := []int{}

	for i := 1 ; i < n + 1 ; i++ {
		q = append(q, i)
	}

	for len(q) > 1 {
		for _ = range k - 1 {
			el := q[0]
			q = q[1:]
			q = append(q, el)
		}
		q = q[1:]
	}

	return q[0]
}

func main() {
	fmt.Println(findWinner(5, 2))
	fmt.Println(findWinner(6, 5))
}
