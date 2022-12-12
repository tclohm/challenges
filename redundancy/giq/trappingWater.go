package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// How much water is trapped
func trapping(heights []int) int {
	var maxLeft, maxRight = 0, 0
	total := 0
	// current height
	for index, _ := range heights {

		maxLeft = 0
		maxRight = 0
		// left side of the current height
		for left := index ; left >= 0 ; left-- {
			// bigger than the maxLeft
			if heights[left] > maxLeft {
				maxLeft = heights[left]
			}
		}
		// right side of current height
		for right := index ; right < len(heights) ; right++ {
			if heights[right] > maxRight {
				maxRight = heights[right]
			}
		}
		// take the minimum of maxleft and right and subtract the current height
		var currentHeight = min(maxLeft, maxRight) - heights[index]
		// add the current height total
		if currentHeight >= 0 {
			total += currentHeight
		} 

	}

	return total
	
}

func opTrap(height []int) int {
	// 1. Id pointer with lesser value
	// 2. Is this pointer lesser than equal to the max of that side
	// 		if true, update max
	// 		if not, get water for pointer and add to total
	// 3. Move pointer inward
	// 4. repeat for other pointer
	var maxLeft, maxRight = 0, 0
	var left, right = 0, len(height) - 1
	var total = 0

	for left < right {
		if height[left] <= height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				total += maxLeft - height[left] 
			}
			left += 1
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				total += maxRight - height[right]
			}
			right -= 1
		}
	}
	return total
}

func main() {
	trapped := []int{0,1,0,2,1,3,1,0,1,2}

	//fmt.Println(trapping(trapped))

	fmt.Println(opTrap(trapped))
}