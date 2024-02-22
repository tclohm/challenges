package main

import "fmt"

func maximalSquare(matrix [][]string) int {
    return 1
}

func main() {
	fmt.Println(maximalSquare([][]string{{"1","0","1","0","0"},{"1","0","1","1","1"},{"1","1","1","1","1"},{"1","0","0","1","0"}}))
	fmt.Println(maximalSquare([][]string{{"0","1"},{"1","0"}}))
	fmt.Println(maximalSquare([][]string{{"0"}}))
}