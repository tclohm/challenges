package main

import (
	"fmt"
	"strconv"
)

func countSeniors(details []string) int {
	count := 0
	for _, person := range details {
		age, err := strconv.Atoi(person[11:13]); if err != nil {
			return -1
		}
		if age > 60 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countSeniors([]string{"7868190130M7522","5303914400F9211","9273338290F4010"}))
	fmt.Println(countSeniors([]string{"1313579440F2036","2921522980M5644"}))
}
