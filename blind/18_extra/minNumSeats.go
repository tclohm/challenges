package main

import (
	"fmt"
	"sort"
	"math"
)

// increase or decrease the position
// return the mininum number of moves for students be in their own seats
func minMoveSeats(seats, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	count := 0

	for i := 0 ; i < len(seats) ; i++ {
		absolute := int(math.Abs(float64(seats[i] - students[i])))
		count += absolute
	}
	return count
}

func main() {
	fmt.Println(minMoveSeats([]int{3,1,5}, []int{2,7,4}))
	fmt.Println(minMoveSeats([]int{4,1,5,9}, []int{1,3,2,6}))
	fmt.Println(minMoveSeats([]int{2,2,6,6}, []int{1,3,2,6}))

}
