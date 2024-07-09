package main

import "fmt"

func avgWaitTime(customers [][]int) int {
	end := 0
	waiting := 0

	for i := 0 ; i < len(customers) ; i++ {
		arrive, length := customers[i][0], customers[i][1]

		if end < arrive {
			end = arrive + length
		} else {
			end = end + length
		}
		waiting += end - arrive
	}

	return waiting / len(customers)
}

func main() {
	fmt.Println(avgWaitTime([][]int{{1,2},{2,5},{4,3}}))
	fmt.Println(avgWaitTime([][]int{{5,2},{5,4},{10,3},{20,1}}))
}
