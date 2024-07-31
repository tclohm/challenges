package main

import "fmt"

func notIn(dst int, set []int) bool {
	for _, s := range set {
		if dst == s {
			return false 
		}
	}
	return true
}

func minCost(source, target int, original, changed, cost []string) int {
	adj := make(map[int][]int)

	for i := 0 ; i < len(original) ; i++ {
		src, dst, currCost := source[i], target[i], cost[i]
		adj[src] = append(adj[src], []int{dst, currCost}...)
 	}

	var dijkstra func (src int) map[int]int
	dijkstra = func (src int) int {
		heap := [][]{{0, src}}
		minCostMap := map[int]int{}

		for len(heap) > 0 {
			// find the smallest, 0 : cost, 1 : node
			sm := []int{1000, 0}
			for _, arr := range heap {
				cost, node := arr[0], arr[1]
				if cost < sm[0] {
					sm[0], sm[1] = cost, node
				}
			}
			cost, node := sm[0], sm[1]

			if _, exist := minCostMap[node]; exist {
				continue
			}

			minCostMap[node] = cost

			for _, neighbor := range adj[node] {
				n, nCost := neighbor[0], neighbor[1]
				heap = append(heap, []int{cost + nCost, n})
			}
			
			return minCostMap
		}
	}

	for _, c := range source {
		minCostMap[c] = dijkstra(c)
	}

	minCostMaps := map[int]int{}
	result := 0

	for i := range source {
		dst := target[i]
		src := source[i]
		if notIn(dst, minCostMap[src]) {
			return -1
		}
		result += minCostMap[src][dst]
	}
	return result
}
