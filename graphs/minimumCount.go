func minIsland(grid [][]int ) int {
	visited := []string{}
	minSize := 10000000000
	for i := 0 ; i < len(grid) ; i++ {
		for j := 0 ; j < len(grid[i]); j++ {
			size := explore(grid, i, j, visited)

			if size > 0 && size < minSize {
				minSize = size
			}
		}
	}
	return minSize
}


func explore(grid [][]int, row int, col int, visited []int) int {
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

	size := 1
	size += explore(grid, row - 1, col, visited)
	size += explore(grid, row + 1, col, visited)
	size += explore(grid, row, col - 1, visited)
	size += explore(grid, row, col + 1, visited)

	return size

}