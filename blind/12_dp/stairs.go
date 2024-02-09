package main

import "fmt"

func climb(steps int) int {
	if steps == 1 {
		return 1
	}
	if steps == 2 {
		return 2
	}
	
	return climb(steps - 1) + climb(steps - 2)
}

func main() {
	fmt.Println(climb(2))
	fmt.Println(climb(3))
	fmt.Println(climb(5))
}