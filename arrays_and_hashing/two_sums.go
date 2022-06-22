package main

import "fmt"

func twoSums(field[]int, target int) []int {
	hm := make(map[int]int)

	for index, number := range field {
		difference := target - number
		position, exists := hm[difference]
		
		if exists {
			return []int{position, index}
		}

		hm[number] = index

	}

	return nil
}

func main() {
	a := []int{2, 7, 11, 16}
	b := []int{3, 2, 4}
	c := []int{3, 3}
	fmt.Println(twoSums(a, 9))
	fmt.Println(twoSums(b, 6))
	fmt.Println(twoSums(c, 6))
}