package main

import "fmt"

func findItinerary(tickets [][]string) []string {
	return []string{}
}

func main() {
	fmt.Println(findItinerary([][]string{{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}}))
	fmt.Println(findItinerary([][]string{{"JFK","SFO"},{"JFK","ATL"},{"SFO","ATL"},{"ATL","JFK"},{"ATL","SFO"}}))
}