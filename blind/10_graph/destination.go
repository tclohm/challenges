package main

import "fmt"

func findDestination(paths [][]string) string {
	ht := map[string]string{}
	// put all the paths into the map
	for _, points := range paths {
		from, to := points[0], points[1]
		ht[from] = to
	}
	// loop through the map 
	// and check if the value, is a key in the map
	// if not that's the end destination
	for _, value := range ht {
		if _, ok := ht[value]; !ok {
			return value
		}
	}

	return "-1"
}

func main() {
	paths := [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
	cities := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "San Paulo"}}

	fmt.Println(findDestination(paths))
	fmt.Println(findDestination(cities))
}