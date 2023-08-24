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

func uniquePath(m, n int) int {
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

func main() {
	fmt.Println(uniquePath(3, 7))
	fmt.Println(uniquePath(3, 2))
}