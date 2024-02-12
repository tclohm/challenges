package main

import "fmt"

func min(a, b int) int {
	if a < b { return a }
	return b
}

// train tickets
func minCost(days, costs []int) int {
	// costs : 1 day = $2 ; 7 days = $7 ; 30 day = $15
	var dfs func (index int) int
	dfs = func (index int) int {
		if index == len(days) { return 0 }

		oneDayTickets := dfs(index + 1) + costs[0]
		
		var i int
		
		for i = index + 1 ; i < len(days) && days[i] < days[index]+7 ; i++ {}
		sevenDaysTicket := dfs(i) + costs[1]

		for i = index + 1 ; i < len(days) && days[i] < days[index]+30 ; i++ {}
		thirtyDaysTicket := dfs(i) + costs[2]

		return min(oneDayTickets, min(sevenDaysTicket, thirtyDaysTicket))
	}

	return dfs(0)
}

func main() {
	// travel on 1, buy 1 day :$2
	// travel 4-8, buy 7 day  :$7
	// travel 20, buy 1 day   :$2
	fmt.Println(minCost([]int{1,4,6,7,8,20}, []int{2,7,15}))
	fmt.Println(minCost([]int{1,2,3,4,5,6,7,8,9,10,30,31}, []int{2,7,15}))
}