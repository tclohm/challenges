package main

import (
	"fmt"
	"math"
)

// there a n cities connected by some number of flights. Given an array flights where 
// flights[i] = [from, to, price]

// given three ints: source, destination, and k (allowed number of stops) 
// FIND THE CHEAPEST PRICE from src to dst

// if there is no route return -1

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    cost := make([]int, n)

    for i := 0 ; i < n ; i++ {
    	cost[i] = math.MaxUint32
    }

   	cost[src] = 0

   	for i := 0 ; i <= k ; i++ {
   		tmp := make([]int, n)
   		copy(tmp, cost)

   		for f := 0 ; f < len(flights) ; f++ {
   			from, to, price := flights[f][0], flights[f][1], flights[f][2]

   			if cost[from] != math.MaxUint32 {
   				tmp[to] = min(tmp[to], cost[from] + price)
   			}
   		}

   		cost = tmp
   	}

   	if cost[dst] == math.MaxUint32 { return -1 }
   	return cost[dst]
}

func min(x, y int) int {
	if x < y { return x }
	return y
}

func main() {
	fmt.Println(findCheapestPrice(4, [][]int{{0,1,100},{1,2,100},{2,0,100},{1,3,600}}, 0, 3, 1))
}