package main

import "fmt"

func twoSum(array []int, target int) []int {
	// key : number, value : position
	var ht = map[int]int{}

	for position, value := range array {
		leftover := target - value
		_, ok := ht[leftover] 
		if !ok {
			ht[value] = position
		} else {
			return []int{ht[leftover],position}
		}
	}

	return nil
}

func main() {
	var array = []int{1,3,7,9,2}
	fmt.Println(twoSum(array, 11))
}