package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func profit(p []int) int {
	var profit int = 0
	var m int = 10000

	for i := 0 ; i < len(p) ; i++ {
		m = min(m, p[i])
		if p[i] - m > profit {
			profit = p[i] - m
		}
	} 
	return profit
}

func profit_w_pointer(p []int) int {
	var left, right = 0, 1
	var max_profit = 0

	for right < len(p) {
		if p[left] < p[right] {
			profit := p[right] - p[left]
			max_profit = max(max_profit, profit) 
		} else {
			left = right
		}
		right += 1
	}
	return max_profit
}

func main() {
	p1 := []int{7,1,5,3,6,4}
	p2 := []int{7,6,4,3,1}
	fmt.Println(profit(p1))
	fmt.Println(profit(p2))
}