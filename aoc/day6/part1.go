package main

import (
	"os"
	"fmt"
	"bufio"
)
// first time a marker could occur is after the 4th char
// marker appears after the seventh char arrives

// we need to identify the first position where the four most recently received characters are all different

// four characters

func main() {
	file, _ := os.Open("file.txt")

	defer file.Close()

	sc := bufio.NewScanner(file)

	var line string

	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "%s", &line)
	}

	for i := 0 ; i < len(line) - 4 ; i++ {
		one := line[i] 
		two := line[i+1]
		three := line[i+2]
		four := line[i+3]

		if one != two && one != three && one != four && two != three && two != four && three != four {
			fmt.Println(i+3+1)
			break
		}
	}
}