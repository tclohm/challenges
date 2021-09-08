package main;
import "fmt";

func depthFirstPrint(graph map[string][]string, source string) {
	// MARK: -- recursion uses a stack don't need to create one. But you can if you want to solve it iteratively
	fmt.Println(source)
	for _, neighbor := range graph[source] {
		depthFirstPrint(graph, neighbor)
	}
}

func main() {
	graph := map[string][]string{
		"a": {"c", "b"},
		"b": {"d"},
		"c": {"e"},
		"d": {"f"},
		"e": {},
		"f": {},
	}

	depthFirstPrint(graph, "a")
}