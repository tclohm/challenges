package main

import "fmt"

func larger(a, b int) int {
	if a > b { return a }
	return b
}

func largestIsland(grid [][]int) int {
	largest := 0

	var dfs func(x, y, current_count int)
	dfs = func(x, y, current_count int) {
		if x < 0 || x >= len(grid[0]) { return }
		if y < 0 || y >= len(grid) { return }
		if grid[y][x] == 0 { 
			largest = larger(largest, current_count)
			return 
		}

		grid[y][x] = 0
		current_count += 1


		// recursive, stacking the stack bruh
		dfs(x + 1, y, current_count)
		dfs(x - 1, y, current_count)
		dfs(x, y + 1, current_count)
		dfs(x, y - 1, current_count)
	}

	for y := 0 ; y < len(grid) ; y++ {
		for x := 0 ; x < len(grid[0]) ; x++ {
			if grid[y][x] == 1 {
				dfs(x, y, 1)
			}
		}
	}

	return largest
}

func main() {
	grid := [][]int{
		{0,0,1,0,0,0,0,1,0,0,0,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,1,1,0,1,0,0,0,0,0,0,0,0},
		{0,1,0,0,1,1,0,0,1,0,1,0,0},
		{0,1,0,0,1,1,0,0,1,1,1,0,0},
		{0,0,0,0,0,0,0,0,0,0,1,0,0},
		{0,0,0,0,0,0,0,1,1,1,0,0,0},
		{0,0,0,0,0,0,0,1,1,0,0,0,0},
	}

	fmt.Println(largestIsland(grid))
}