package main

import "fmt"

// 0 empty
// 1 fresh
// 2 rotten
func rotting(grid [][]int) int {
	q := [][]int{}
	fresh := 0
	time := 0
	for r := 0 ; r < len(grid) ; r++ {
		for c := 0 ; c < len(grid[0]) ; c++ {
			if grid[r][c] == 1 {
				fresh += 1
			}
			if grid[r][c] == 2 {
				q = append(q, []int{r, c})
			}
		}
	}

	directions := [4][2]int{{0,1},{0,-1},{1,0},{-1,0}}

	for fresh > 0 && len(q) != 0 {
		for i := 0 ; i < len(q) ; i++ {

			item := q[0]
			q = q[1:]

			for _, dir := range directions {
				row, col := item[0] + dir[0], item[1] + dir[1]

				if row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0]) && grid[row][col] == 1 {
					grid[row][col] = 2
					q = append(q, []int{row, col})
					fresh -= 1
				}
			}
		}
		time += 1
	}

	if fresh == 0 {
		return time
	}

	return -1
}

func main() {
	fmt.Println(rotting([][]int{{2,1,1},{1,1,0},{0,1,1}}))
}