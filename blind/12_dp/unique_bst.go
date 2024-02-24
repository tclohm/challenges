package main

import "fmt"
// 1 => 1
// 2 => 2
// 3 => 5
// 4 => 
func numTrees(n int) int {
	G := make([]int, n + 1)
	G[0] = 1
	G[1] = 1

	for i := 2 ; i <= n ; i++ {
		for j := 1 ; j <= i ; j++ {
			G[i] += G[j - 1] * G[i - j]
		}
	}
	return G[n]
}

func recursiveNumTrees(n int) int {
	var dfs func(k int) int
	dfs = func(k int) int {
		var result = 0
		if k == 1 || k == 0 {
			return 1
		}
		for i := 0 ; i < k ; i++ {
			result += dfs(i) * dfs(k - i - 1)
		}
		return result
	}
	return dfs(n)
}

func main() {
	fmt.Println(recursiveNumTrees(1))
	fmt.Println(recursiveNumTrees(2))
	fmt.Println(recursiveNumTrees(3))
	fmt.Println(recursiveNumTrees(4))
	fmt.Println(recursiveNumTrees(5))
}