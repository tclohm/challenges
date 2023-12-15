package main

import "fmt"

func findDestination(paths [][]string) string {
	ht := map[string]string{}

	for _, points := range paths {
		from, to := points[0], points[1]
		ht[from] = to
	}

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