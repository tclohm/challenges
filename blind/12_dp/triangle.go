package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	// go left or right
	var dfs func (x, y int) int
	dfs = func (x, y int) int {
		if y >= len(triangle) { return 0 }
		if x > len(triangle[y]) { return 0 }

		min := min(triangle[y][x], triangle[y][x + 1])

		if min == triangle[y][x] {
			return min + dfs(x, y+1)
		}

		return min + dfs(x+1, y+1)
	}
	return triangle[0][0] + dfs(0,1)
}

func main() {
	fmt.Println(minimumTotal([][]int{{2},{3,4},{6,5,7},{4,1,8,3}}))
	fmt.Println(minimumTotal([][]int{{-10}}))
}