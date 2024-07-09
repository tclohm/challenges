package main

import "fmt"
func findCriticalPoints(nums []int) []int {
	var critical func (previous, current, next int)
	critical = func (previous, current, next int) {
		return (
			previous > current < next || previous < current > next
		)
	}

	var (
		previous = 0
		current = 1
		next = 2
		min_distance = 100
		max_distance = -100

		previous_critical_index = 0
		first_critical_index = 0

		index = 1
	)

	for next < len(nums) {
		previous, current, next = current, next, next + 1
		index += 1
	}
	return [min_distance, max_distance]
}

func main() {
	fmt.Println(findCriticalPoints([]int{3,1}))	
	fmt.Println(findCriticalPoints([]int{5,3,1,2,5,1,2}))
}
