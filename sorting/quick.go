package main

import "fmt"

func quick(elements []int, below int, upper int) {
	if below < upper {
		var part int = divide(elements, below, upper)
		quick(elements, below, part-1)
		quick(elements, part+1, upper)
	}
}

func divide(elements []int, below int, upper int) int {
	var center int = elements[upper]
	var i int = below

	for j := below ; j < upper ; j++ {
		if elements[j] <= center {
			swap(&elements[i], &elements[j])
			i += 1
		}
	}

	swap(&elements[i], &elements[upper])
	return i 
}

func swap(element1 *int, element2 *int) {
	var tmp int = *element1
	*element1 = *element2
	*element2 = tmp
}

func main() {
	var num int

	fmt.Print("Enter Number of elements: ")
	fmt.Scan(&num)

	var array = make([]int, num)

	for i := 0 ; i < num ; i++ {
		fmt.Print("array[", i, "]: ")
		fmt.Scan(&array[i])
	}

	fmt.Print("Elements: ", array, "\n")
	quick(array, 0, num-1)
	fmt.Print("Sorted:", array, "\n")
}