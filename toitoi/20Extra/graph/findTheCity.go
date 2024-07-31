package main

import (
	"fmt"
	"math"
	"container/heap"
)

type Edge struct {
	node, cost int
}
// Heap
type PriorityQueue []Edge

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Edge)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// greedy algo
func dijkstra(cityInfo map[int][]Edge, idx, n, distanceAllowed int) int {
	temp := make([]int, n)
	for i := range temp {
		temp[i] = math.MaxInt32
	}
	temp[idx] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Edge{idx, 0})

	for pq.Len() > 0 {
		popped := heap.Pop(pq).(Edge)
		cost, node := popped.cost, popped.node
		// go through neighbors
		for _, neighbor := range cityInfo[node] {
			pathSum := cost + neighbor.cost
			if temp[neighbor.node] > pathSum {
				temp[neighbor.node] = pathSum
				heap.Push(pq, Edge{neighbor.node, pathSum})
			}
		}
	}

	freq := 0
	for _, d := range temp {
		if d <= distanceAllowed {
			freq++
		}
	}
	return freq
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	// Adjacency list
	cityInfo := make(map[int][]Edge)
	for _, it := range edges {
		v1, v2, weight := it[0], it[1], it[2]
		cityInfo[v1] = append(cityInfo[v1], Edge{v2, weight})
		cityInfo[v2] = append(cityInfo[v2], Edge{v1, weight})
	}
	// big value initiated
	minCount := math.MaxInt32
	result := -1
	// src node...ascending order
	for i := 0; i < n; i++ {
		count := dijkstra(cityInfo, i, n, distanceThreshold)
		if count <= minCount {
			minCount = count
			result = i
		}
	}

	return result
}


func main() {
	fmt.Println(findTheCity(4, [][]int{{0,1,3},{1,2,1},{1,3,4},{2,3,1}}, 4))
} 
