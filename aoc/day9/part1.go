package main

import (
	"fmt"
	"os"
	"bufio"
)

// rope with knot at each end
// if the end is ever two steps directly up, down, left, or right from the tail,
// the tail must alos move one step in that direction so it remains close

// if the head and tail aren't touching and aren't in the same row or column,
// the tail always moves one step diagonally to keep up

// find where the tail goes as the head follows a series of motions
// head and tail start at the same position

// how many positions does the tail visit at least once?
type point struct {
	x, y int
}

func adjust(tail, head point) (t point) {
	t = tail
	switch (point{head.x-tail.x, head.y-tail.y}) {
	case point{-2, 1}, point{-1, 2}, point{0, 2}, point{1, 2}, point{2, 1}:
		t.y++
	}
	switch (point{head.x-tail.x, head.y-tail.y}) {
	case point{1, 2}, point{2, 1}, point{2, 0}, point{2, -1}, point{1, -2}:
		t.x++
	}
	switch (point{head.x-tail.x, head.y-tail.y}) {
	case point{2, -1}, point{1, -2}, point{0, -2}, point{-1, -2}, point{-2, -1}:
		t.y--
	}
	switch (point{head.x-tail.x, head.y-tail.y}) {
	case point{-1, -2}, point{-2, -1}, point{-2, 0}, point{-2, 1}, point{-1, 2}:
		t.x--
	}
	return
}

func main() {
	file, _ := os.Open("file.txt")

	defer file.Close()

	sc := bufio.NewScanner(file)

	var direction string
	var steps int

	var head = point{0, 0}
	var tail = point{0, 0}

	var visited = make(map[point]bool)

	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "%s %d", &direction, &steps)

		for steps > 0 {
			switch direction {
				case "U":
					head.y++
				case "R":
					head.x++
				case "D":
					head.y--
				case "L":
					head.x--
			}
			steps--
			tail = adjust(tail, head)
			visited[tail] = true
		}
	}

	fmt.Println(len(visited))

}