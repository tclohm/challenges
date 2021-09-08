package main;
import "fmt"

func is(node string, in map[string][]string) bool {
	for key, _ := range in {
		if key == node {
			return true
		}
	}
	return false
}

func in(node string, in []string) bool {
	for _, key := range in {
		if key == node {
			return true
		}
	}
	return false
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

func hasPath(graph map[string][]string, source string, destination string, visited []string) bool {
	if source == destination {
		return true
	}

	if in(source, visited) {
		return false
	}

	visited = append(visited, source)

	for _, neighbor := range graph[source] {
		if hasPath(graph, neighbor, destination, visited) == true {
			return true
		}
	}

	return false
}


func main() {
	edges := [][]string{{"i", "j"}, {"k", "i"}, {"m", "k"}, {"k", "l"}, {"o", "n"}}
	graph := buildGraph(edges)
	visited := []string{}
	fmt.Println(hasPath(graph, "n", "i", visited)) // false
	fmt.Println(hasPath(graph, "j", "m", visited)) // true
	fmt.Println(hasPath(graph, "l", "i", visited)) // true
}