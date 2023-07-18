package main

import "fmt"

func numberOfIslands(grid [][]string) int {
	islands := 0
	var backtrack func(x, y int)
	backtrack = func(x, y int) {
		// check bounds
		if x < 0 || x >= len(grid) { return }
		if y < 0 || y >= len(grid[0]) { return }
		if grid[x][y] == "0" { return }

		grid[x][y] = "0"

		backtrack(x - 1, y)
		backtrack(x + 1, y)
		backtrack(x, y - 1)
		backtrack(x, y + 1)
	}

	for row := 0 ; row < len(grid) ; row++ {
		for col := 0 ; col < len(grid[0]) ; col++ {
			if grid[row][col] == "1" {
				islands += 1
				// row -- y, col -- x
				backtrack(col, row)
			}
		}
	}
	return islands
}

func main() {
	grid := [][]string{
		{"1","1","1","1","0"},
		{"1","1","0","1","0"},
		{"1","1","0","0","0"},
		{"0","0","0","0","0"},
	}

	grid2 := [][]string{
  		{"1","1","0","0","0"},
  		{"1","1","0","0","0"},
  		{"0","0","1","0","0"},
  		{"0","0","0","1","1"},
	}

	fmt.Println(numberOfIslands(grid))
	fmt.Println(numberOfIslands(grid2))
}