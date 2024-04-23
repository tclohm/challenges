package main

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	adjacency := make(map[int][]int)
	for i := range n {
		adjacency[i] = []int{}
	}
	
	for _, nodes := range edges {
		n1, n2 := nodes[0], nodes[1]
		adjacency[n1] = append(adjacency[n1], n2)
		adjacency[n2] = append(adjacency[n2], n1)
	}

	edge_count := map[int]int{}
	leaves := []int{}
	for src, neighbors := range adjacency {
		if len(neighbors) == 1 {
			leaves = append(leaves, src)
		}
		edge_count[src] = len(neighbors)
	}

	for len(leaves) > 0 {
		if n <= 2 {
			return leaves
		}
		snapShotLength := len(leaves)
		for i := 0 ; i < snapShotLength; i++ {

			node := leaves[0]
			leaves = leaves[1:]
			n -= 1
			for _, neighbor := range adjacency[node] {
				edge_count[neighbor] -= 1
				if edge_count[neighbor] == 1 {
					leaves = append(leaves, neighbor)
				}
			}
		}
	}

	return []int{0}
}

func main() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1,0},{1,2},{1,3}}))
	fmt.Println(findMinHeightTrees(6, [][]int{{3,0},{3,1},{3,2},{3,4},{5,4}}))
}