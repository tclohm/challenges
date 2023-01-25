package main

import (
	"os"
	"bufio"
	"fmt"
)

func increment(cycleNumber, x *int) {
	if *cycleNumber % 40 == 0 && *cycleNumber <= 220 {
		fmt.Println()
	}
	if *x - 1 == *cycleNumber % 40 || *x == *cycleNumber % 40 || *x + 1 == *cycleNumber % 40 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	*cycleNumber++
}

func main() {
	file, _ := os.Open("file.txt")

	defer file.Close()

	sc := bufio.NewScanner(file)

	x, cycle := 1, 0

	var direction string
	var number int

	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "%s %d", &direction, &number)
		increment(&cycle, &x)
		if direction == "addx" {
			increment(&cycle, &x)
			x += number
		}
	}
}