package main

import "fmt"

// minutes to be used to not be grumpy
func satisfied(customers, grumpy []int, minutes int) int {
	var (
		window = 0
		maxWindow = 0
		satisfied = 0
		l = 0
	)
	
	for r, _ := range customers {
		// if grumpy, increase window
		if grumpy[r] == 1 {
			window += customers[r]
		} else {
			satisfied += customers[r]
		}
		// check window if minutes is less
		if r - l + 1 > minutes {
			if grumpy[l] == 1 {
				window -= customers[l]
			}
			l += 1
		}
		// check max window
		maxWindow = max(window, maxWindow)
	}
	return satisfied + maxWindow
}

func main() {
	fmt.Println(satisfied([]int{1,0,1,2,1,1,7,5}, []int{0,1,0,1,0,1,0,1}, 3))
	fmt.Println(satisfied([]int{1}, []int{0}, 1))
}
