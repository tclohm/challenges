package main

import "fmt"

// There is a robot on an m x n grid. 
//The robot is initially located at the top-left corner 
// (i.e., grid[0][0]). 
// The robot tries to move to the bottom-right corner 
// (i.e., grid[m - 1][n - 1]). 
// The robot can only move either down or right at any point in time.

// Given the two integers m and n, 
// return the number of possible unique paths 
// that the robot can take to reach the bottom-right corner.

func uniquePathDFS(m, n int) int {
	var explore func(i, j int) int
	explore = func(i, j int) int {
		if i == m && j == n {
			return 1
		}

		var ipaths, jpaths int

		if i < m {
			ipaths = explore(i + 1, j)
		}

		if j < n {
			jpaths = explore(i, j + 1)
		}

		return ipaths + jpaths
	}
	return explore(1, 1)
}

func uniquePathMemo(m, n int) int {
	paths := make([][]int, n)
	for i := 0 ; i < len(paths) ; i++ {
		paths[i] = make([]int, m)
	}
	paths[n - 1][m - 1] = 1

	var explore func(i, j int) int
	explore = func(i, j int) int {
		if paths[j][i] != 0 {
			return paths[j][i]
		}

		var ipaths, jpaths int

		if i < m - 1 {
			ipaths = explore(i + 1, j)
		}
		if j < n - 1 {
			jpaths = explore(i, j + 1)
		}
		paths[j][i] = ipaths + jpaths
		return ipaths + jpaths
	}

	return explore(0, 0)
}

func uniquePath(m, n int) int {
	paths := make([][]int, n)

	for i := 0 ; i < len(paths) ; i++ {
		paths[i] = make([]int, m)
	}

	for i := 0 ; i < m ; i++ {
		paths[0][i] = 1
	}

	for i := 0 ; i < n ; i++ {
		paths[i][0] = 1
	}

	for i := 1 ; i < m ; i++ {
		for k := 1; k < n ; k++ {
			paths[k][i] = paths[k][i - 1] + paths[k - 1][i]
		}
	}

	return paths[n - 1][m - 1]
}

func main() {
	fmt.Println(uniquePath(3, 7))
	fmt.Println(uniquePath(3, 2))
}