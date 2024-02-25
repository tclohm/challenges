package main

import "fmt"
// 1 => 1
// 2 => 2
// 3 => 5
// 4 => 
// T: O(n**2) S O(n)
func numTrees(n int) int {
	result := make([]int, n + 1)
	result[0] = 1
	result[1] = 1

	for i := 2 ; i <= n ; i++ {
		for j := 1 ; j <= i ; j++ {
			result[i] += result[j - 1] * result[i - j]
		}
	}
	return result[n]
}

// Catalan number
func recursiveNumTrees(n int) int {
	var dfs func(k int) int
	dfs = func(k int) int {
		var result = 0
		if k == 1 || k == 0 {
			return 1
		}
		for i := 0 ; i < k ; i++ {
			// for given root
			// result += left side * right side
			result += dfs(i) * dfs(k - i - 1)
		}
		return result
	}
	return dfs(n)
}

func main() {
	fmt.Println(numTrees(1))
	fmt.Println(numTrees(2))
	fmt.Println(numTrees(3))
	fmt.Println(numTrees(4))
	fmt.Println(numTrees(5))
}