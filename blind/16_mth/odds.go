package main

import "fmt"

func count(low, high int) int {
	return (( high + 1 ) / 2) - (low / 2)
}

func main() {
	fmt.Println(count(2,7)) // 3
	fmt.Println(count(3,7)) // 3
	fmt.Println("---")
	fmt.Println(count(4,7)) // 2
	fmt.Println(count(5,7)) // 2
	fmt.Println("---")
	fmt.Println(count(6,7)) // 1
	fmt.Println(count(7,7)) // 1
}