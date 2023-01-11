package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)

// one item per line
// each elf separated their own inventory by a blank line

// which elf to ask, 
// who is caring the most calories


func main() {
	input, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	maxCalories := 0
	currentCalories := 0

	for scanner.Scan() {
		snack, err := strconv.Atoi(scanner.Text())
		currentCalories += snack

		if err != nil {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
		}
	}
	fmt.Println(maxCalories)
}