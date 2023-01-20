package main

import (
	"fmt"
	"bufio"
	"os"
)

// 0, short - 9, tallest

// left, up, right, down

// how many trees are visible from outside the grid

// we can see all the exterior trees

func main() {
	file, _ := os.Open("file.txt")

	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		fmt.Println(sc.Text())
	}
}