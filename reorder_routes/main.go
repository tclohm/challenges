package main

import "fmt"

// 	 -> 1 -> 3 <- 2
// 0
//   <- 4 -> 5

type Edge struct {
	a, b int
}

type Set struct {
	container map[Edge]struct{}
}

func NewSet() *Set {
	s := &Set{
		container: make(map[Edge]struct{}),
	}
	return s
}

func (s *Set) Exists(key Edge) bool {
	_, exists := s.container[key]
	return exists
}

func (s *Set) Add(key Edge) {
	s.container[key] = struct{}{}
}

func minReorder(n int, connections [][]int) int {
	// start at city 0
	// recursively check it's neighbors
	// count outgoing edges
	var edges = NewSet()
	var neighbors = make(map[int][]int)

	var visits = map[int]bool{}

	var changes int = 0
	var cp *int = &changes

	// fill up set with all edges
	for index := 0 ; index < len(connections) ; index++ {
		a, b := connections[index][0], connections[index][1]
		e := Edge{a, b}
		edges.Add(e)
	}

	for city := 0 ; city < n ; city++ {
		neighbors[city] = []int{}
	}

	// find the neighbors
	for index := 0 ; index < len(connections) ; index++ {
		a, b := connections[index][0], connections[index][1]
		neighbors[a] = append(neighbors[a], b)
		neighbors[b] = append(neighbors[b], a)
	}

	fmt.Println("edges", edges.container)
	fmt.Println("v", visits, "c", changes)
	fmt.Println("neigh", neighbors)

	var dfs func(int, Set, map[int][]int, map[int]bool, *int)
	dfs = func(city int, edges Set, neighbors map[int][]int, visits map[int]bool, cp *int) {
		for _, neighbor := range neighbors[city] {
			_, exists := visits[neighbor]; if exists {
				continue
			}
			e := Edge{neighbor, city}
			ok := edges.Exists(e); if !ok {
				*cp += 1
			}
			fmt.Println(*cp, &cp)
			visits[neighbor] = true
			dfs(neighbor, edges, neighbors, visits, cp) 
		}
	}
	visits[0] = true
	dfs(0, *edges, neighbors, visits, cp)
	
	return changes
}

func main() {
	var connections = [][]int {{0,1},{1,3},{2,3},{4,0},{4,5}}
	fmt.Println(minReorder(6, connections))
}