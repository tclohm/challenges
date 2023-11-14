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

type times [][]int

func dijkstra(times times, n, k int) int {
	type edge struct { to, weight int }

	w := make(map[int][]edge)
	for _, time := range times {
		from, to, weight := time[0], time[1], time[2]
		w[from] = append(w[from], edge{to, weight})
	}

	dist := make([]int, n+1)
	for i := 1 ; i <= n ; i++ { dist[i] = -1 }

	intree := make([]bool, n+1)

	intree[k] = true
	dist[k] = 0

	v := k
	for {
		intree[v] = true

		for _, e := range w[v] {
			if dist[e.to] == -1 || (dist[e.to] > dist[v] + e.weight) {
				dist[e.to] = dist[v] + e.weight
			}
		}

		minv, mind := -1, -1
		for i := 1 ; i <= n ; i++ {
			if !intree[i] && (mind == -1 || dist[i] < mind) {
				minv, mind = i, dist[i]
			}
		}

		if minv == -1 { break }

		v = minv
	}

	max := dist[1]
	for i := 1 ; i <= n ; i++ {
		if dist[i] == -1 { return -1 }
		if dist[i] > max { max = dist[i] }
	}

	return max
}

func bellmanFord(times times, n, k int) int {
	dist := make([]int, n+1)
	for i := 1 ; i <= n ; i++ { dist[i] = -1 }

	dist[k] = 0

	for i := 0 ; i < n - 1 ; i++ {
		for _, time := range times {
			from, to, weight := time[0], time[1], time[2]

			if dist[from] == -1 { continue }

			if dist[to] == 1 || dist[from] + weight < dist[to] {
				dist[to] = dist[from] + weight
			}
		}
	}

	max := dist[1]

	for i := 1 ; i <= n ; i++ {
		if dist[i] == -1 { return -1 }
		if dist[i] > max { max = dist[i] }
	}

	return max
}

func main() {

	fmt.Println(networkDelay([][]int{{2,1,1},{2,3,1},{3,4,1}}, 4, 2)) // 2
	fmt.Println(networkDelay([][]int{{1,2,1}}, 2, 1)) // 1
	fmt.Println(networkDelay([][]int{{1,2,1}}, 2, 2)) // -1

}