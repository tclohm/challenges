package main

import "fmt"


func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num & 1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}

func main() {
	
	var n uint32 = 00000000000000000000000000001011
	var n2 uint32 = 00000000000000000000000010000000
	var n3 uint32 = 00000000000000000000011111111111

	fmt.Println(hammingWeight(n))
	fmt.Println(hammingWeight(n2))
	fmt.Println(hammingWeight(n3))
}