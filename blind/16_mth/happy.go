package main

import (
	"fmt"
	"strconv"
)

func happy(n int) bool {
	// happy number
	// starts with a positive integer, replace the number by the sum of the 
	// squares of its digits
	// Repeat the process until number = 1
	// end in 1 are happy
	total := 0
	seen := make(map[int]bool)
	for {
		if seen[n] { break }
		seen[n] = true

		number := fmt.Sprint(n)

		length := len(number)

		for i := 0 ; i < length ; i++ {
			digit, err := strconv.Atoi(string(number[i]))
			if err != nil {
				fmt.Println("error")
				return false
			}
			total += (digit * digit)
		}

		if total == 1 { 
			return true
		}

		n = total

		total = 0
	}

	return false
}

func main() {
	fmt.Println(happy(19))
	fmt.Println(happy(2))
}