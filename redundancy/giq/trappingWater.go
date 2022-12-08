package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func trapping(water []int) int {
	var maxLeft, maxRight = 0, 0
	total := 0

	for index, _ := range water {

		maxLeft = 0
		maxRight = 0

		for left := index ; left >= 0 ; left-- {
			if water[left] > maxLeft {
				maxLeft = water[left]
			}
		}

		for right := index ; right < len(water) ; right++ {
			if water[right] > maxRight {
				maxRight = water[right]
			}
		}

		var currentWater = min(maxLeft, maxRight) - water[index]

		if currentWater >= 0 {
			total += currentWater
		} 

	}

	return total
	
}

func main() {
	trapped := []int{0,1,0,2,1,3,1,0,1,2}

	fmt.Println(trapping(trapped))
}