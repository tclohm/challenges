package main

import "fmt"

// [7,1,5,3,6,4]


func aSolution(prices []int) int {
	biggestProfit := 0
	for buy_pointer := 0 ; buy_pointer < len(prices) - 1 ; buy_pointer++ {
		sell_pointer := buy_pointer + 1
		
		if sell_pointer == len(prices) { return biggestProfit }

		end := len(prices)
		// --MARK: sliding window
		for sell_pointer < end {

			if prices[sell_pointer] - prices[buy_pointer] > biggestProfit {
				biggestProfit = prices[sell_pointer] - prices[buy_pointer]
			}

			sell_pointer++
		}

	}
	return biggestProfit
}

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