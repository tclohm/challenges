package main

import (
	"fmt"
	"sort"
)

func sortPeople(names []string, heights []int) []string {

	var heightToNames = map[int]string{}
	var result = []string{}
	for index, name := range names {

		heightToNames[heights[index]] = name

	}

	sort.Ints(heights)

	for index := len(heights) - 1 ; index >= 0 ; index-- {
		name := heightToNames[heights[index]]
		result = append(result, name)
	}
	return result

} 

func main() {
	fmt.Println(sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))
	fmt.Println(sortPeople([]string{"Alice", "Bob", "Bob"}, []int{155, 185, 150}))
}
