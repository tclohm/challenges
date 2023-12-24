package main

import "fmt"

func generate(rows int) [][]int {
	var result = [][]int{}
	for i := 0 ; i < rows ; i++ {
		level := make([]int, i+1, i+1)
		level[0] = 1
		level[len(level) - 1] = 1
		result = append(result, level)
	}

	if len(result) < 2 { return result }

	for row := 0 ; row < rows ; row++ {
		for column := 0 ; column < len(result[row]) ; column++ {
			if result[row][column] == 0 {
				result[row][column] = result[row-1][column-1] + result[row-1][column]
			}
		}
	}

	return result
}

func main() {
	fmt.Println(generate(1))
	fmt.Println(generate(2))
	fmt.Println(generate(5))
}