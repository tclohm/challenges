package main;
import "fmt";


func breadthFirstPrint(graph map[string][]string, source string) {
	var queue = []string{source}
	for len(queue) != 0 {
		var current string = queue[0]
		queue = queue[1:]
		fmt.Println(current)
		for _, neighbor := range graph[current] {
			queue = append(queue, neighbor)
		}
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

	breadthFirstPrint(graph, "a")
}