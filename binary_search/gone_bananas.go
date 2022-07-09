package main

import (
	"fmt"
	"math"
)

// Time Complexity O(nm)
// Space O(1)
func goneBananasBruteForce(piles []int, h int) int {
	// start at speed = 1
	// current speed, calculate how many hours monkey needs to eat all of the piles
	//	if monkey cannot finish piles within `h` hours, increment speed by 1
	//	speed += 1 and start over
	//	if monkey finishes all piles within `h` hours, return speed
	var speed int = 1
	for {
		var hoursSpent int = 0

		for _, pile := range piles {
			hoursSpent += int(math.Ceil(float64(pile) / float64(speed)))
		}

		if hoursSpent <= h {
			return speed
		} else {
			speed += 1
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bsEatingSpeed(piles []int, h int) int {
	var left int = 1
	var right int = 0

	for _, pile := range piles {
		right = max(pile, right)
	}

	for left < right {
		middle := left + (right - left) / 2
		hoursSpent := 0

		for _, pile := range piles {
			hoursSpent += int(math.Ceil(float64(pile) / float64(middle)))
		}

		if hoursSpent <= h {
			right = middle
		} else {
			left = middle + 1
		}
	}
	return right
}

func main() {
	p1 := []int{3,6,7,11}
	
	fmt.Println(bsEatingSpeed(p1, 8))

	p2 := []int{30,11,23,4,20}

	fmt.Println(bsEatingSpeed(p2, 5))

	p3 := []int{30,11,23,4,20}

	fmt.Println(bsEatingSpeed(p3, 6))
}