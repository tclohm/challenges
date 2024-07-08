package main

import "fmt"

// exchange empty water bottles for full ones
// return the max number of water bottles we can drink
func howMuch(have int, exchangeRate int) int {
	// total := have
	// for i := 1 ; i < total ; i += exchangeRate {
	// 	total++
	// }
	// return total
	return have + (have - 1) / (exchangeRate - 1)
}

func main() {
	fmt.Println(howMuch(9, 3))
	fmt.Println(howMuch(15,4))
}
