package main

import "fmt"

func reduncancy(edges [][]int) []int {
    parent := make([]int, len(edges) + 1)

    for connection := 0 ; connection < len(parent) ; connection++ {
    	parent[connection] = connection
    }

    for _, edge := range edges {
    	if find(parent, edge[0]) == find(parent, edge[1]) {
    		return edge
    	}

    	unify(parent, edge[0], edge[1])
    }

    return []int{}
}

func find(parent []int, num int) int {
	if parent[num] == num {
		return num
	}

	return find(parent, parent[num])
}


func unify(parent []int, x, y int) {
	parent[find(parent, y)] = find(parent, x)
}

func main() {
	fmt.Println(reduncancy([][]int{{1,2},{1,3},{2,3}}))
	fmt.Println(reduncancy([][]int{{1,2},{2,3},{3,4},{1,4},{1,5}}))
}