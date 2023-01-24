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

	file, _ := os.Open("file.txt")

	defer file.Close()

	sc := bufio.NewScanner(file)

	forest := [][]rune{}
	//treeHouseViews := [][]int{}

	for sc.Scan() {
		row := []rune{}

		for _, tree := range sc.Text() {
			row = append(row, tree)
		}
		forest = append(forest, row)

		//treeHouseViews = append(treeHouseViews, make([]int, len(sc.Text()), len(sc.Text())))
	}

	bestView := -1

	for row := 0 ; row < len(forest) ; row++ {
		for col := 0 ; col < len(forest[0]) ; col++ {

			currentTreeHeight := forest[row][col]

			// look 'up' until we hit edge or a tree that is taller than our potential tree house
			outlook := row - 1
			northView := 1
			southView := 1
			westView := 1
			eastView := 1

			if outlook < 0 { northView = 0 }
			for outlook > 0 {
				northTreeHeight := forest[outlook][col]
				if northTreeHeight < currentTreeHeight {
					outlook -= 1
					northView += 1
				} else {
					break
				}
			}

			// look 'right' until we hit edge or a tree that is taller than our potential tree house

			outlook = col + 1

			if outlook >= len(forest[0]) { eastView = 0 }
			for outlook < len(forest[0]) - 1 {
				eastTreeHeight := forest[row][outlook]
				if eastTreeHeight < currentTreeHeight {
					outlook += 1
					eastView += 1
				} else {
					break
				}
			}

			// look 'down' until we hit edge or a tree that is taller than our potential tree house

			outlook = row + 1

			if outlook >= len(forest) { southView = 0 }
			for outlook < len(forest) - 1 {
				southTreeHeight := forest[outlook][col]
				if southTreeHeight < currentTreeHeight {
					outlook += 1
					southView += 1
				} else {
					break
				}
			}
			
			// look 'left' until we hit edge or a tree that is taller than our potential tree house

			outlook = col - 1

			if outlook < 0 { westView = 0 }
			for outlook > 0 {
				westTreeHeight := forest[row][outlook]
				if westTreeHeight < currentTreeHeight {
					outlook -= 1
					westView += 1
				} else {
					break
				}
			}

			view := northView * eastView * southView * westView

			//treeHouseViews[row][col] = view

			if view > bestView {
				bestView = view
			}
		}
	}

	fmt.Println(bestView)
}