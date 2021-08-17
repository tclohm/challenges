func swapLexOrder(str string, pairs [][]int) string {

	edges := make(map[int][]int)
	for _, pair := range pairs {
		edges[pair[0]-1] = append(edges[pair[0]-1], pair[1]-1)
		edges[pair[1]-1] = append(edges[pair[1]-1], pair[0]-1)
	}

	runes := []rune(str)

	for idx := range runes {
		pos := FindBestSwapPosition(idx, runes, edges)
		runes[idx], runes[pos] = runes[pos], runes[idx]

	}

	return string(runes)
}

func FindBestSwapPosition(idx int, runes []rune, edges map[int][]int) int {
	best := idx

	// perform bfs
	visited := make(map[int]bool)
	tovisit := make([]int, 1)
	tovisit[0] = idx

	for {
		if len(tovisit) <= 0 {
			break
		}

		// pop
		var i int
		i, tovisit = tovisit[0], tovisit[1:]
		visited[i] = true

		if i > idx && runes[i] > runes[best] {
			best = i
		}

		oute, exists := edges[i]
		if !exists {
			break
		}

		for _, e := range oute {
			_, seen := visited[e]
			if !seen {
				tovisit = append(tovisit, e)
			}
		}
	}

	return best
}
