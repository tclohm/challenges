package main

import "fmt"

// [7,1,5,3,6,4]

func maxProfit(prices []int) int {
	var profit = 0
	var minPrice = prices[0]

	for i := 1 ; i < len(prices) ; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
		
		if (prices[i] - minPrice) > profit {
			profit = prices[i] - minPrice
		}
	}

	return profit
}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
}