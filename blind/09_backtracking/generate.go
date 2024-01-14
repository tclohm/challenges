package main

import "fmt"

func generate(n int) []string {
	result := []string{}

	var dfs func(n, openCount, closeCount int, current string, array *[]string)
	dfs = func(n int, openCount, closeCount int, current string, array *[]string) {
		if openCount == n && closeCount == n {
			*array = append(*array, current)
			return
		}

		if openCount < n {
			dfs(n, openCount + 1, closeCount, current + "(", array)
		}

		if openCount > closeCount {
			dfs(n, openCount, closeCount + 1, current + ")", array)
		} 
	}

	dfs(n, 0, 0, "", &result)

	return result
}

func main() {
	fmt.Println(generate(3))
	fmt.Println(generate(1))
}