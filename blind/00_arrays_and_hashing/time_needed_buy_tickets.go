package main

import "fmt"

func timeNeeded(tickets []int, k int) int {
	time := 0

	for i, v := range tickets {
		// have not passed our kth index
		if i <= k {
			time += min(v, tickets[k])
		} else {
			// we passed our kth index, means our ticket[k] has gone down one
			time += min(v, tickets[k] - 1)
		}
	}
	return time
}

// slower
func timeRequired(tickets []int, k int) int {
	time := 0
	index := 0

	for tickets[k] > 0 {
		if tickets[i] > 0 {
			tickets[i]--
			time++
		}
		i++

		if i >= len(tickets) {
			i = 0
		}
	}
	
	return time
}

func main() {
	fmt.Println(timeNeeded([]int{2,3,2}, 2))
	fmt.Println(timeNeeded([]int{5,1,1,1}, 0))
	fmt.Println(timeNeeded([]int{1,2,5,3,10,5,7}, 3))
}