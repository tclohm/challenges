package main

import (
	"fmt"
	"math"
)


func maximumDetonation(bombs [][]int) int {

	adjacency := make(map[int][]int)

    for i, bomb := range bombs {
    	adjacency[i] = bomb
    }


    for i := 0 ; i < len(bombs) ; i++ {
    	for j := i + 1 ; j < len(bombs) ; j++ {
    		x1, y1, r1 := bombs[i][0], bombs[i][1], bombs[i][2]
    		x2, y2, r2 := bombs[j][0], bombs[j][1], bombs[j][2]
    		diameter := int(math.Sqrt(float64((x1 - x2) * (x1 - x2)) + float64((y1 - y2) * (y1 - y2))))

    		if diameter <= r1 {
    			adjacency[i] = append(adjacency[i], j)
    		}
    		if diameter <= r2 {
    			adjacency[j] = append(adjacency[j], i)
    		}
    	}
    }

    var dfs func(i int, visited map[int]bool) int
    dfs = func(i int, visited map[int]bool) int {
    	if _, ok :=  visited[i]; ok {
    		return 0
    	}
    	visited[i] = true
    	for _, neighbor := range adjacency[i] {
    		dfs(neighbor, visited)
    	}
    	return len(visited)
    }

    result := 0
    for i := range bombs {
    	result = max(result, dfs(i, map[int]bool{}))
    }

    return result
}

func main() {
	fmt.Println(maximumDetonation([][]int{{2,1,3},{6,1,4}}))
	fmt.Println(maximumDetonation([][]int{{1,1,5},{10,10,5}}))
	fmt.Println(maximumDetonation([][]int{{1,2,3},{2,3,1},{3,4,2},{4,5,3},{5,6,4}}))
}