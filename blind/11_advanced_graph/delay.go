package main

import "fmt"

// give an network of n nodes, 1 to n
// given times, list of travel times as directed edges
// times[i] = (src, dst, time it take to receive signal)

// send a signal from a given node k
// return the min time it takes for all the n nodes
// to receive the signal


// visit each node and store the shortest distance from k to the node
// avoid circular paths, dont visit a node twice if the current distance from
// k to the node is greater than one obtained earlier

// whole graph is traversed, return the max shortest distance

func networkDelay(times [][]int, n, k int) int {
	children := make(map[int][][2]int{})
	for _, time := range times {
		children[time[0]] = append(children[time[0]], [2]int{time[1], time[2]})
	}
	seen := map[int]bool{}
	minPath := map[int]int{}

	var dfs func(int, int)
	dfs = func(node, weight int) {
		seen[node] = true
		if _, ok := minPath[node]; !ok {
			minPath[node] = weight
		} else if weight >= minPath[node] {
			return
		} else {
			minPath[node] = weight
		}

		for _, child := range children[node] {
			dfs(child[0], weight+child[1])
		}
	}

	dfs(k, 0)

	if len(seen) != n { return -1 }

	maxDist := -1

	for _, distance := range minPath {
		maxDist = max(maxDist, distance)
	}

	return maxDist
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func main() {

	fmt.Println(networkDelay([][]int{{2,1,1},{2,3,1},{3,4,1}}, 4, 2)) // 2
	fmt.Println(networkDelay([][]int{{1,2,1}}, 2, 1)) // 1
	fmt.Println(networkDelay([][]int{{1,2,1}}, 2, 2)) // -1

}