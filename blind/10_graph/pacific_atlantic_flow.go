package main

import "fmt"

// island gets rained on, and the water can flow to neighbors (n,w,s,e)
// neighbors must be less than or equal to the current cell's height
// return array of coords that allow water to reach both oceans

func island(heights [][]int) [][]int {
	ROWS, COLS := len(heights), len(heights[0])
	pacific, atlantic := make(map[int]bool), make(map[int]bool)

	var dfs func(int, int, int, map[int]bool)
	dfs = func (row, col, previousHeight int, visited map[int]bool) {
		if visited[row * COLS + col] || row < 0 || row == ROWS || col < 0 || col == ROWS || heights[row][col] < previousHeight { return }
		visited[row * COLS + col] = true

		dfs(row + 1,     col, heights[row][col], visited)
		dfs(row - 1,     col, heights[row][col], visited)
		dfs(row    , col + 1, heights[row][col], visited)
		dfs(row    , col - 1, heights[row][col], visited)
	}

	for col := 0 ; col < COLS ; col++ {
		dfs(0       , col, heights[0][col], pacific)
		dfs(ROWS - 1, col, heights[ROWS - 1][col], atlantic)
	}

	for row := 0 ; row < ROWS ; row++ {
		dfs(row, 0       , heights[row][0], pacific)
		dfs(row, COLS - 1, heights[row][COLS - 1], atlantic)
	}

	result := make([][]int, 0)
	for r := 0 ; r < ROWS ; r++ {
		for c := 0 ; c < COLS ; c++ {
			if pacific[r * COLS + c] && atlantic[r * COLS + c] {
				result = append(result, []int{r, c})
			}
		}
	}
	return result
}

func main() {
	fmt.Println(island([][]int{{1,2,2,3,5},{3,2,3,4,4},{2,4,5,3,1},{6,7,1,4,5},{5,1,1,2,4}}))
}