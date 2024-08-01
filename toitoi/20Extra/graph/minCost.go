package main

import (
	"fmt"
	"container/heap"
	"math"
)


func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	graph := make([][]int, 26)
	for i := range original {
		index := original[i] - 'a'
		if graph[index] == nil {
			graph[index] = make([]int, 26)
		}

		if graph[original[i]-'a'][changed[i]-'a'] == 0 {
			graph[original[i]-'a'][changed[i]-'a'] = cost[i]
		} else {
			graph[original[i]-'a'][changed[i]-'a'] = min(cost[i], graph[original[i]-'a'][changed[i]-'a'])
		}
	}

	cache := make([][]int, 26)
	for i := range cache {
		cache[i] = make([]int, 26)
	}

	res := 0
	for i := range source {
		if source[i] == target[i] {
			continue
		}

		if cache[source[i]-'a'][target[i]-'a'] > 0 {
			res += cache[source[i]-'a'][target[i]-'a']
			continue
		}

		changeCost := djikstra(graph, source[i], target[i], &cache)
		if changeCost == -1 {
			return -1
		}

		if cache[source[i]-'a'][target[i]-'a'] == 0 {
			cache[source[i]-'a'][target[i]-'a'] = changeCost
		} else {
			cache[source[i]-'a'][target[i]-'a'] = min(cache[source[i]-'a'][target[i]-'a'], changeCost)
		}

		res += cache[source[i]-'a'][target[i]-'a']
	}

	return int64(res)
}

func djikstra(graph [][]int, source, target byte, cache *[][]int) int {
	pq := &PQ{}
	heap.Init(pq)

	heap.Push(pq, Edge{source, 0})
	visited := make([]bool, 26)

	for pq.Len() > 0 {
		root := heap.Pop(pq).(Edge)

		if visited[root.char-'a'] {
			continue
		}

		visited[root.char-'a'] = true

		(*cache)[source-'a'][root.char-'a'] = root.cost

		if root.char == target {
			return root.cost
		}

		thresholdCost := math.MaxInt32
		if len(graph[root.char-'a']) != 0 {
			thresholdCost = graph[root.char-'a'][target-'a']
			if thresholdCost == 0 {
				thresholdCost = math.MaxInt32
			} else {
				thresholdCost += root.cost
			}
		}

		for i, edgeCost := range graph[root.char-'a'] {
			cost := root.cost + edgeCost
			if edgeCost == 0 || thresholdCost < cost {
				continue
			}

			heap.Push(pq, Edge{byte('a' + i), cost})
		}
	}

	return -1
}

type Edge struct {
	char byte
	cost int
}

type PQ []Edge

func (p PQ) Len() int {
	return len(p)
}

func (p PQ) Less(i, j int) bool {
	return p[i].cost < p[j].cost
}

func (p PQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PQ) Push(x any) {
	*p = append(*p, x.(Edge))
}

func (p *PQ) Pop() any {
	res := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return res
}

func main() {
	fmt.Println(minimumCost("abcd", "acbe", []byte{'a','b','c','c','e','d'}, []byte{'b','c','b','e','b','e'}, []int{2,5,5,1,2,20}))
	fmt.Println(minimumCost("aaaa", "bbbb", []byte{'a','c'}, []byte{'c','b'}, []int{1,2}))
}
