package main

import (
	"fmt"
	"os"
	"bufio"
)

func isContained(startOne, endOne, startTwo, endTwo int) bool {

	// check if s1 and e1 contain s2 and e2
	if startTwo >= startOne && endTwo <= endOne || startOne >= startTwo && endOne <= endTwo {
		return true
	}

	return false
}

// how many assignment pairs does one range full container the other

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
	}

	points := 0

	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		var startF, endF, startS, endS int
		fmt.Sscanf(sc.Text(), "%d-%d,%d-%d", &startF, &endF, &startS, &endS)

		if isContained(startF, endF, startS, endS) {
			points++
		}

	}

	fmt.Println(points)
}