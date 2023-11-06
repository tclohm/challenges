package main

import (
	"fmt"
	"strconv"
)

func multiply(n1, n2 string) string {
	first, err := strconv.Atoi(n1)
	if err != nil { return "error" }

	second, err := strconv.Atoi(n2)
	if err != nil { return "error" }

	return fmt.Sprint(first * second)
}

func main() {
	fmt.Println(multiply("2", "3"))
	fmt.Println(multiply("123", "456"))
}