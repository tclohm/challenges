package main

import (
	"fmt"
	"math/rand"
	"time"
)

func create(num int) []int {
	var array []int
	array = make([]int, num, num)
	rand.Seed(time.Now().UnixNano())

	for i := 0 ; i < num ; i++ {
		array[i] = rand.Intn(99999) - rand.Intn(99999)
	}
	return array
}

func sort(array []int) []int {

	if len(array) < 2 {
		return array
	}

	var middle int = len(array) / 2

	return merge(sort(array[:middle]), sort(array[middle:]))
}


func merge(left []int, right []int) []int {
	var num int
	var i int
	var j int

	num, i, j = len(left) + len(right), 0, 0

	var array []int = make([]int, num, num)

	for k := 0 ; k < num ; k++ {
		if i > len(left) - 1 && j <= len(right) - 1 {
			array[k] = right[j]
			j++
		} else if j > len(right) - 1 && i <= len(left) - 1 {
			array[k] = left[i]
			i++
		} else if left[i] < right[j] {
			array[k] = left[i]
			i++
		} else {
			array[k] = right[j]
			j++
		}
	}
	return array
}

func main() {
	var elements []int = create(40)
	fmt.Println("\n Before sorted \n\n", elements)
	fmt.Println("\n After sorting\n\n", sort(elements), "\n")
}