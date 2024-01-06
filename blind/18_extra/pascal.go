package main

import "fmt"

func generate(rows int) [][]int {
	if rows == 0 { return [][]int{} }
	var container = [][]int{{1}}
	if rows == 1 { return container }
	for i := 1 ; i < rows ; i++ {
		container = append(container, make([]int, i+1, i+1))
		end := len(container[i]) - 1
		container[i][0] = 1
		container[i][end] = 1
	}

	for r := 1 ; r < len(container) ; r++ {
		for c := 1 ; c < len(container[r]) - 1 ; c++ {
			left := c - 1
			right := c
			above := container[r - 1]

			if left < 0 || right > len(above) {
				continue
			} else {
				container[r][c] = container[r - 1][left] + container[r - 1][right]
			}
		}
	}

	return container
}

func main() {
	fmt.Println(generate(5))
	fmt.Println(generate(1))
	fmt.Println(generate(6))
}