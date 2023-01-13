package main

import (
	"fmt"
	"bufio"
	"os"
)

// how many overlap

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}

	points := 0

	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		var firstStart, firstEnd, secondStart, secondEnd int
		fmt.Sscanf(sc.Text(), "%d-%d,%d-%d", &firstStart, &firstEnd, &secondStart, &secondEnd)
		
		if secondStart <= firstEnd && secondEnd >= firstStart || firstStart <= secondEnd && firstEnd >= secondStart {
			points += 1
		}

	}

	fmt.Println(points)
}