package main

import "fmt"

func majority(elements []int) int {
	majority_element, majority_element_freq := elements[0], 1
	for _, element := range elements {
		if majority_element_freq == 0 {

			majority_element, majority_element_freq = element, 1

		} else {

			if element == majority_element {
				majority_element_freq += 1
			} else {
				majority_element_freq -= 1
			}

		}
	}

	return majority_element

}

func main() {
	fmt.Println(majority([]int{3,2,3}))
	fmt.Println(majority([]int{2,2,1,1,1,2,2}))
}