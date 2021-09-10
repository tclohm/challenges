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

// func exploreSize(graph map[int][]int, current int, visited []int) int {
// 	if has(visited, current) {
// 		return 0
// 	}

// 	visited = append(visited, current)
// 	_ = copy()

// 	size := 1

// 	for _, neighbor := range graph[current] {
// 		fmt.Println(visited)
// 		size += exploreSize(graph, neighbor, visited)
// 	}

// 	return size
// }

// func largestComponent(graph map[int][]int) int {
// 	largest := 0
// 	visited := []int{}
// 	for node, _ := range graph {
// 		size := exploreSize(graph, node, visited)
// 		if size > largest {
// 			largest = size
// 		}
// 	}
// 	return largest
//}

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
	largest := largestComponent(graph)
	fmt.Println(largest)
}