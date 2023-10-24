package main

import "fmt"

func sum(a, b int) int {
	for b != 0 {
		tmp := (a & b) << 1
		a = (a ^ b)
		b = tmp
	}
	return a
}

func main() {
	fmt.Println(sum(1,2))
	fmt.Println(sum(2,23))
}