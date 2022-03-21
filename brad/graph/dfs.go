package main

import "fmt"

type Graph map[string][]string
type Node string


func traverse(node string, counter []int, visited map[string]bool, traversal_times *map[string]map[string]int, graph Graph) {

	if _, ok := visited[node]; !ok {
		visited[node] = true
	}

	counter[0] += 1
	
	if (*traversal_times)[node] == nil {
		(*traversal_times)[node] = map[string]int{}
	}
	(*traversal_times)[node]["discovery"] = counter[0]

	for _, neighbor := range graph[node] {
		if _, ok := visited[neighbor]; !ok {
			traverse(neighbor, counter, visited, traversal_times, graph)
		}
	}

	counter[0] += 1

	(*traversal_times)[node]["finish"] = counter[0]
}


func dfs(graph Graph, start string) map[string]map[string]int {
	visited := map[string]bool{}
	counter := []int{0}

	traversal_times := map[string]map[string]int{}

	traverse(start, counter, visited, &traversal_times, graph)

	return traversal_times
}

func main() {

	simple_graph := map[string][]string{
		"A": {"B", "D"},
		"B": {"C", "D"},
		"C": {},
		"D": {"E"},
		"E": {"B", "F"},
		"F": {"C"},
	}

	tt := dfs(simple_graph, "A")

	fmt.Println(tt)
}