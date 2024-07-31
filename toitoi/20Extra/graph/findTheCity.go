package main

import (
	"fmt"
	"math"
	"container/heap"
)

type Edge struct {
	node, cost int
}

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

func comeOnDijkstra(cityInfo map[int][]Edge, idx, n, distanceAllowed int) int {
	temp := make([]int, n)
	for i := range temp {
		temp[i] = math.MaxInt32
	}
	temp[idx] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Edge{idx, 0})

	for pq.Len() > 0 {
		costVertex := heap.Pop(pq).(Edge)
		cost, vertex := costVertex.cost, costVertex.node

		for _, it := range cityInfo[vertex] {
			pathSum := cost + it.cost
			if temp[it.node] > pathSum {
				temp[it.node] = pathSum
				heap.Push(pq, Edge{it.node, pathSum})
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
	cityInfo := make(map[int][]Edge)
	for _, it := range edges {
		cityInfo[it[0]] = append(cityInfo[it[0]], Edge{it[1], it[2]})
		cityInfo[it[1]] = append(cityInfo[it[1]], Edge{it[0], it[2]})
	}
	ans := math.MaxInt32
	result := -1

	for i := 0; i < n; i++ {
		path := comeOnDijkstra(cityInfo, i, n, distanceThreshold)
		if path <= ans {
			ans = path
			result = i
		}
	}

	return result
}


func main() {
	fmt.Println(findTheCity(4, [][]int{{0,1,3},{1,2,1},{1,3,4},{2,3,1}}, 4))
} 
