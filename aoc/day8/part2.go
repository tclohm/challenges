package main

import (
	"fmt"
	"os"
	"bufio"
)

// want to see a lot of trees

// measure the viewing distance from a given tree : look up, down, left, and right
// stop if you reach an edge or at the first tree that is the same height or taller than the tree under construction
// if the tree is at the edge, the tree will have at least one of the viewing distances at zero


// tree : look left, look right, look up, look down
// viewing distance : left trees seen * up trees seen * right trees seen * down trees seen

// what is the highest viewing distance
func main() {

	file, _ := os.Open("test.txt")

	defer file.Close()

	sc := bufio.NewScanner(file)

	forest := [][]rune{}
	views := [][]int{}

	for sc.Scan() {
		row := []rune{}

		for _, tree := range sc.Text() {
			row = append(row, tree)
		}
		forest = append(forest, row)
		views = append(views, make([]int, len(sc.Text()), len(sc.Text())))

	}

	fmt.Println(forest, views)

}