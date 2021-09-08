package main;
import "fmt";

func hasPath(graph map[string][]string, source string, destination string) bool {
	var queue = []string{ source }
	var set = []string{}
	//fmt.Println("---- start", source, "----")
	for len(queue) != 0 {

		current := queue[0]
		queue = queue[1:]

		//fmt.Println("current node:", current)
		//fmt.Println("queue:", queue)
		
		for _, visited := range set {
			if current == visited {
				return false
			}
		}

		set = append(set, current)

		if current == destination {
			//fmt.Println("---- reached", destination, "----")
			return true
		}

		for _, neighbor := range graph[current] {
			queue = append(queue, neighbor)
		}
	}
	return false
}

func hasPathDFS(graph map[string][]string, source, destination string) bool {
	if source == destination {
		return true
	}

	for _, neighbor := range graph[source] {
		if (hasPath(graph, neighbor, destination) == true) {
			return true
		}
	}

	return false
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

	resultOne 	:= hasPath(graph, "a", "b")
	resultTwo 	:= hasPath(graph, "a", "f")
	resultThree := hasPath(graph, "c", "a")
	resultFour 	:= hasPath(graph, "f", "e")
	resultFive 	:= hasPath(graph, "d", "f")
	fmt.Println("bfs", resultOne)
	fmt.Println("bfs", resultTwo)
	fmt.Println("bfs", resultThree)
	fmt.Println("bfs", resultFour)
	fmt.Println("bfs", resultFive)

	fmt.Println("---- dfs has path ----")

	dfsResultOne 	:= hasPathDFS(graph, "a", "b")
	dfsResultTwo 	:= hasPathDFS(graph, "a", "f")
	dfsResultThree 	:= hasPathDFS(graph, "c", "a")
	dfsResultFour 	:= hasPathDFS(graph, "f", "e")
	dfsResultFive 	:= hasPathDFS(graph, "d", "f")

	fmt.Println("dfs", dfsResultOne)
	fmt.Println("dfs", dfsResultTwo)
	fmt.Println("dfs", dfsResultThree)
	fmt.Println("dfs", dfsResultFour)
	fmt.Println("dfs", dfsResultFive)
}