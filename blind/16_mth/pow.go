package main

import "fmt"

func pow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}
	pow := 1.0
	for n != 0 {
		if n % 2 == 1 {
			pow *= x
		}
		x *= x
		n /= 2
	}
	return pow
}

func main() {
	fmt.Println(pow(2.000, 10))
	fmt.Println(pow(2.010, 3))
	fmt.Println(pow(2.000, -2))
}