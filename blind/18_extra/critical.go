package main

import "fmt"
func findCriticalPoints(nums []int) []int {
	var critical func (previous, current, next int) bool
	critical = func (previous, current, next int) bool {
		return (previous > current && current < next) || (previous < current && current > next)
	}

	var (
		previous = 0
		current = 1
		next = 2
		min_distance = 100
		max_distance = -1

		previous_critical_index = 0
		first_critical_index = 0

		index = 1
	)

	for next < len(nums) {
		if critical(previous, current, next) {
			if first_critical_index != 0 {
				max_distance = index - first_critical_index
				min_distance = min(min_distance, index - previous_critical_index)
			} else {
				first_critical_index = index
			}
			previous_critical_index = index
		}
		previous, current, next = current, next, next + 1
		index += 1
	}

	if min_distance == 100 {
		min_distance = -1
	}

	return []int{min_distance, max_distance}
}

func main() {
	fmt.Println(findCriticalPoints([]int{3,1}))	
	fmt.Println(findCriticalPoints([]int{5,3,1,2,5,1,2}))
}
