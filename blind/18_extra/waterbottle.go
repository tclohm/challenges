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
	// have + (have - 1) / (exchangeRate - 1)

	drank := 0
	empty := 0

	for have > 0 {
		drank += have
		empty += have
		have = empty / exchangeRate
		empty = empty % exchangeRate
	}
	return drank
}

func main() {
	fmt.Println(howMuch(9, 3))
	fmt.Println(howMuch(15,4))
}
