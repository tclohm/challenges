package main

import "fmt"

func reverse(n uint32) uint32 {
	var res uint32 = 0

	for i := 0 ; i < 32 ; i++ {
		bit := (n >> i) & 1
		res = res | (bit << (31 - i))
	}

	return res
}

func main() {
	fmt.Println(reverse(00000000000000000000001010011100))
	fmt.Println(reverse(00000000000000000000000001111101))
}