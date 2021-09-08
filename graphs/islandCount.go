func islandCount(grid [][]int) int {
	visited := []string{}
	count := 0

	for i := 0 ; i < len(grid) ; i++ {
		for j := 0 ; j < len(grid[i]); j++ {
			if explore(grid, i, j, visited) {
				count++
			}
		}
	}
	return count
}

func explore(grid [][]int, row int, col int, visited []int) bool {
	rowBounds := 0 <= row && r < len(grid)
	colBounds := 0 <= col && col < len(grid)
	if !rowBounds || !colBounds {
		return false
	}

	if grid[row][col] == "W" {
		return false
	}

	position := row + "," + col
	if exists(visited, position) {
		return false
	}
	visited = append(visited, position)

	explore(grid, row - 1, col, visited)
	explore(grid, row + 1, col, visited)
	explore(grid, row, col - 1, visited)
	explore(grid, row, col + 1, visited)

	return true

}