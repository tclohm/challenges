// n = number of nodes in the graph
// g = adjaceny list rep unweighted graph
func bfs(start int, end int) []int {
	prev = solve(start)
	return reconstructPath(start, end, prev)
} 

func solve(start) []int {
	q := make([]int, 0, 100)
	append(q, start)
	visited := make([]bool, 0, 100)
	visited[start] = true
	prev := make([]int, 0, 100)

	for len(q) {
		node := q[0]
		q = q[1:]
		
		neighbors := g.get(node)

		for next := 0 ; next < len(neighbors) ; next++ {
			if !visited[next] {
				q = append(q, next)
				visited[next] = true
				prev[next] = node
			}
		}
	}
	return prev
}

func reconstructPath(s, e, prev) {
	path := make([]int, 0, 100)
	for at := e ; at != nil ; at = prev[at] {
		path = append(path, at)
	}

	for i, j := 0, len(path) - 1 ; i < j ; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i] 
	}

	if path[0] == s {
		return path
	}
	return path
}