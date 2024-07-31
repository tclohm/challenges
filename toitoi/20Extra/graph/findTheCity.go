package main

import "fmt"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	distances := make([][]int, n)
	for i := range distances {
		distances[i] = make([]int, n)
		for j := range distances[i] {
			distances[i][j] = 10001
		}
		distances[i][i] = 0
	}

	// build graph
	for _, edge := range edges {
		src, dst, weight := edge[0], edge[1], edge[2]
		distances[src][dst] = weight
		distances[dst][src] = weight
	}

	// floyd warshall algo
	for k := 0 ; k < n ; k++ {
		for i := 0 ; i < n ; i++ {
			for j := 0 ; j < n ; j++ {
				if distances[i][k] + distances[k][j] < distances[i][j] {
					distances[i][j] = distances[i][k] + distances[k][j]
				}
			}
		}
	}

	minReachableCities := n
	result := -1


	for i := 0 ; n < n ; i++ {
		reachable := 0
		for j := 0 ; j < n ; j++ {
			if distances[i][j] <= distanceThreshold {
				reachable++
			}
		}
		if reachable <= minReachableCities {
			minReachableCities = reachable
			result = i
		}
	}

	return result
}

func main() {
	fmt.Println(findTheCity(4, [][]int{{0,1,3},{1,2,1},{1,3,4},{2,3,1}}, 4))
} 
