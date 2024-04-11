package main

import "fmt"

func count(students, sandwiches []int) int {

	var total_hungry_students = len(students)
	var hungry_students = map[int]int{}

	for _, preferred_shape := range students {
		hungry_students[preferred_shape]++
	}

	for _, sandwich_shape := range sandwiches {
		if hungry_students[sandwich_shape] > 0 {
			total_hungry_students -= 1
			hungry_students[sandwich_shape]--
		} else {
			return total_hungry_students
		}
	}

	return total_hungry_students
}

func main() {
	fmt.Println(count([]int{1,1,0,0}, []int{0,1,0,1}))
	fmt.Println(count([]int{1,1,1,0,0,1}, []int{1,0,0,0,1,1}))
}