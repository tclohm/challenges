package main

import "fmt"

func generate(rows int) [][]int {
	var result = [][]int{}
	for i := 0 ; i < rows ; i++ {
		level := make([]int, i+1, i+1)
		result = append(result, level)
	}
	return result
}

func main() {
	fmt.Println(generate(1))
	fmt.Println(generate(5))
}