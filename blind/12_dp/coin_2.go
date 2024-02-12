package main

import "fmt"

// return the number of combinations that make up this amount
func change(amount int, coins []int) int {
	var dfs func (current_amount int, index int) int
	dfs = func (current_amount int, index int) int {
		if current_amount == amount { return 1 }
		if current_amount > amount { return 0 }
		if index == len(coins) { return 0 }
		return dfs(current_amount + coins[index], index) + dfs(current_amount, index + 1)
	}

	return dfs(0, 0)
}

func main() {
	fmt.Println(change(5, []int{1,2,5}))
	fmt.Println(change(3, []int{2}))
	fmt.Println(change(10, []int{10}))
}