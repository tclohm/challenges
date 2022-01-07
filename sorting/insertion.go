package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomSequence(num int) []int {
	var sequence []int = make([]int, num, num)
	rand.Seed(time.Now().UnixNano())

	for i := 0 ; i < num ; i++ {
		sequence[i] = rand.Intn(999) - rand.Intn(999)
	}

	return sequence
}

func insertionSort(elements []int) {
	var length int = len(elements)

	for i := 1 ; i < length ; i++ {
		var j int = i
		for j > 0 {
			if elements[j - 1] > elements[j] {
				elements[j - 1], elements[j] = elements[j], elements[j - 1]
			}
			j = j - 1
		}
	}
}

func main() {
	var sequence []int = randomSequence(24)
	fmt.Println("\n---- Before Sorting ---- \n\n", sequence)
	insertionSort(sequence)
	fmt.Println("\n---- After Sorting ---- \n\n", sequence, "\n")
}