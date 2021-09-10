package main;
import "fmt";

// have to look at the details of golang :) --- NOT working as expected. visited keeps reseting to an empty array

// func has(array []int, node int) bool {
// 	for _, value := range array {
// 		if value == node {
// 			return true
// 		}
// 	}
// 	return false
// }

// func explore(graph map[int][]int, current int, visited []int) bool {
// 	if has(visited, current) {
// 		return false
// 	}

// 	visited = append(visited, current)

// 	for _, neighbor := range graph[current] {
// 		explore(graph, neighbor, visited)
// 	}

// 	return true
// }

// func connectedComponentCount(graph map[int][]int) int {
// 	count := 0
// 	visited := []int{}
// 	queue := []
// 	for node, _ := range graph {
// 		if explore(graph, node, visited) {
// 			count++
// 		}
// 	}
// 	return count
// }
func in(node int, arr *[]int) bool {
	for _, n := range *arr {
		if node == n {
			return true
		}
	}
	return false
}

func exist(node int, arr **[]int) bool {
	for _, n := range **arr {
		if node == n {
			return true
		}
	}
	return false
}

func dfs(graph map[int][]int, node int, visited *[]int) int {

	stack := []int{node}

	for len(stack) != 0 {
		fmt.Println("stack", stack)
		length := len(stack) - 1
		n := stack[length]
		stack = stack[:length]

		if exist(n, &visited) {
			fmt.Println("node", n, "in", *visited)
			continue
		}

		fmt.Println("node", n, "not in visited", *visited)

		*visited = append(*visited, n)

		for _, neighbor := range graph[n] {
			stack = append(stack, neighbor)
		}
	}

	return 1
}

func connectedComponentCount(graph map[int][]int) int {
	visited := []int{}
	count := 0

	for node, _ := range graph {
		if in(node, &visited) {
			fmt.Println("node", node, "in", "visited", visited)
			continue
		}
		number := dfs(graph, node, &visited)
		count += number
	}

	return count
}

func main() {
	graph := map[int][]int{
		0: {8, 1, 5},
		1: {0},
		5: {0, 8},
		8: {0, 5},
		2: {3, 4},
		3: {2, 4},
		4: {3, 2},
	}
	count := connectedComponentCount(graph)
	fmt.Println(count)
}