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

func eff(students, seats []int) int {
	maxIndex := max(maxIn(students), maxIn(seats)) + 1
	countSeats := make([]int, maxIndex, maxIndex)
	countStudents := make([]int, maxIndex, maxIndex)

	for _, seat := range seats {
		countSeats[seat] += 1
	}
	for _, student := range students {
		countStudents[student] += 1
	}

	i, j := 0, 0
	res := 0
	remain := len(students)
	for remain > 0 {
		if countSeats[i] == 0 {
			i += 1
		}
		if countStudents[j] == 0 {
			j += 1
		}
		if countSeats[i] > 0 && countStudents[j] > 0 {
			res += int(math.Abs(float64(i - j)))
			countSeats[i] -= 1
			countStudents[j] -= 1
			remain -= 1
		}
	}
	return res
}

func maxIn(nums []int) int {
	big := -1
	for _, n := range nums {
		if n > big {
			big = n
		}
	}
	return big
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(eff([]int{3,1,5}, []int{2,7,4}))
	fmt.Println(eff([]int{4,1,5,9}, []int{1,3,2,6}))
	fmt.Println(eff([]int{2,2,6,6}, []int{1,3,2,6}))

}
