package main

import "fmt"
// bottom up -- what's the minimum amount of coins does it take to get to 0
// now what's the minimum amount of coins to get to 1, and then 2, then 3, then 4
// use the knowledge of previous calcuations
func change(coins []int, amount int) int {
	// dp
	changes := make([]int, amount+1)
	// make all the change in the array the amount + 1
	for i := range changes {
		changes[i] = amount + 1
	}
	// first change is 0
	changes[0] = 0
	// now iterate through the amount with a nested for loop
	for i := 1 ; i < amount + 1 ; i++ {
		for _, c := range coins {
			if c <= i {
				changes[i] = min(changes[i], 1 + changes[i - c])
			}
		}
	}

	if changes[amount] != amount + 1 {
		return changes[amount]
	}

	return -1
} 

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(change([]int{1,2,5}, 11))
	fmt.Println(change([]int{2}, 3))
	fmt.Println(change([]int{1}, 0))
}