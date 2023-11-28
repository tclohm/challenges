package main

type Node struct {
	Value int
}

type adjlist [][]int

func largestPath(colors string, edges adjlist) int {

	for src, dst := range edges {
		adj[src] = append(adj[src], dst)
	}

	var dfs func (node *Node)
	dfs = func (node *Node) {
		if _, ok := path[node]; ok {
			return 1000000
		}
		if _, ok := visited[node]; ok {
			return 0
		}

		visited = append(visited, node)
		path = append(path, node)

		colorIndex := colors[node] - 'a'

		for _, neighbor := range adj[node] {
			if dfs(neighbor) == 1000000 {
				return 1000000
			}

			for _, color := range(26) {
				var v int
				if color == colorIndex {
					v = 1
				} else {
					v = 0
				}
				count[node][color] = max(count[node][color], v + count[neighbor][color])
			}

			remove(path, node)
			return max(count[node])
		}
	}

	n, res := len(colors), 0
	visited, path := []*Node{}, []*Node{}
	count := make([]int, 0, 26)
	
	for i := 0 ; i < n ; i++ {
		res = max(dfs(i), res)
	}
	if res == 1000000 {
		return -1
	}
	return res
	return -1
}