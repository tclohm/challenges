package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
)

type Calories []int
func (c Calories) Len() int { return len(c) }
func (c Calories) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c Calories) Less(i, j int) bool { return c[i] < c[j] } 

func main() {
	input, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}

	var array Calories = []int{}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	currentCalories := 0

	for scanner.Scan() {
		snack, err := strconv.Atoi(scanner.Text())
		currentCalories += snack

		if err != nil {
			array = append(array, currentCalories)
			currentCalories = 0
		}
	}

	sort.Sort(array)

	total := 0

	for i := 1 ; i < 4 ; i++ {
		fmt.Println(i)
		total += array[len(array) - i]
	}

	fmt.Println(total)
}