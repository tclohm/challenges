package main;
import "fmt";

func shortestPath(edges [][]int, start int, end int) int {
	graph := buildGraph(edges)
	visited := []int{start}
	queue := [][]int{{start,0}}

	for queue != 0 {
		current := queue[0]
		queue = queue[1:]

		node := current[0]
		distance := current[1]

		if (node == end) {
			return distance
		}

		for _, neighbor := range graph[node] {
			if !in(visited, neighbor) {
				visited = append(visited, neighbor)
				queue = append({neighbor, distance + 1})
			}
		}
	}
	return -1
}

func buildGraph(edges [][]string) map[string][]string {
	var graph = make(map[string][]string)

	for _, edge := range edges {

		a, b := edge[0], edge[1]

		if !is(a, graph) {
			graph[a] = []string{}
		}
		if !is(b, graph) {
			graph[b] = []string{}
		}

		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)

	}

	return graph
}