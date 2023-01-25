package main

import (
	"os"
	"bufio"
	"fmt"
)

// signal being sent by cpu

// starts with 1

// two instructions
// addx V -- two cycles to complete, two cycles x register is increased by the value V
// noop -- one cycle to complete, no othe effect

// small program

// noop
// addx 3
// addx -5

// X = 1
// addx = 4
// addx = -1

// signal strength -- cycle multiplied by the value of the X register during the 20th cycle

func increment(cycleNumber, x, total *int) {
	*cycleNumber++
	if (*cycleNumber - 20) % 40 == 0 && *cycleNumber <= 220 {
		*total += *x * *cycleNumber
	}
}

func main() {
	file, _ := os.Open("file.txt")

	defer file.Close()

	sc := bufio.NewScanner(file)

	x, cycle := 1, 0

	var direction string
	var number int
	var value int

	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "%s %d", &direction, &number)
		increment(&cycle, &x, &value)
		if direction == "addx" {
			increment(&cycle, &x, &value)
			x += number
		}
	}

	fmt.Println(value)
}