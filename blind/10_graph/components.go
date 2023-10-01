package main

import "fmt"

func count(n int, edges [][]int) int {
	if n == 1 { return 1 }

	var components int = 0

	var visited = make([]bool, n)
	var graph = make(map[int][]int)

	for _, edge := range edges {
		var to = edge[0]
		var from = edge[1]

		graph[to] = append(graph[to], from)
		graph[from] = append(graph[from], to)
	}

	var dfs func (int)

	dfs = func (node int) {
		visited[node] = true

		neighbors := graph[node]

		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	for node := 0 ; node < n ; node++ {
		if !visited[node] {
			components++
			dfs(node)
		}
	}

	return components
}

func main() {
	fmt.Println(count(3, [][]int{{0,1},{0,2}}))
	fmt.Println(count(6, [][]int{{0,1},{1,2},{2,3},{4,5}}))
}