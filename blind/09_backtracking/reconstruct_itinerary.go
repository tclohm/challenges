package main

import (
	"fmt"
	"sort"
)

func findItinerary(tickets [][]string) []string {

	adjacent_list := map[string][]string{}


	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		adjacent_list[src] = append(adjacent_list[src], dst)
	}

	for _, ticket := range tickets {
		src := ticket[0]
		sort.Strings(adjacent_list[src])
	}

	var result = []string{"JFK"}

	var backtrack func(src string) bool
	backtrack = func(src string) bool {
		if len(result) == len(tickets) + 1 {
			return true
		}
		if len(adjacent_list[src]) == 0 {
			return false
		}

		temporary := append([]string{}, adjacent_list[src]...)

		for i, v := range temporary {
			popped := adjacent_list[src][i]
			adjacent_list[src] = append(adjacent_list[src][:i], adjacent_list[src][i+1:]...)
			
			result = append(result, v)

			if backtrack(v) { 
				return true 
			}
			if len(adjacent_list[src]) == i {
				adjacent_list[src] = append(adjacent_list[src], popped)
			} else {
				adjacent_list[src] = append(adjacent_list[src][:i+1], adjacent_list[src][i:]...)
				adjacent_list[src][i] = popped
			}
			
		}
		return false
	}

	backtrack("JFK")

	return result
}

func main() {
	fmt.Println(findItinerary([][]string{{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}}))
	fmt.Println(findItinerary([][]string{{"JFK","SFO"},{"JFK","ATL"},{"SFO","ATL"},{"ATL","JFK"},{"ATL","SFO"}}))
}