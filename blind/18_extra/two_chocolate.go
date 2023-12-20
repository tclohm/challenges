package main

import "fmt"

func buy(prices []int, money int) int {
	min1, min2 := 10000, 10000

	for _, price := range prices {
		if price < min1 {
			min1, min2 = price, min1
		} else if price < min2 {
			min2 = price
		}
	}

	leftover := money - min1 - min2

	if leftover >= 0 {
		return leftover
	}

	return money
}

func main() {
	fmt.Println(buy([]int{1,2,2}, 3))
	fmt.Println(buy([]int{3,2,3}, 3))
}