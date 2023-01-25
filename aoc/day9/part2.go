package main

import (
	"os"
	"fmt"
	"bufio"
)

type point struct {
	x, y int
}

func adjust(tail, head point) (t point) {
	t = tail
	switch (point{head.x-tail.x,head.y-tail.y}){
	case point{-2,1},point{-1,2}, point{0,2}, point{1,2}, point{2,1}, point{2,2}, point{-2,2}:
		t.y++
	}
	switch (point{head.x-tail.x,head.y-tail.y}){
	case point{1,2},point{2,1}, point{2,0}, point{2,-1}, point{1,-2}, point{2,2}, point{2,-2}:
		t.x++
	}
	switch (point{head.x-tail.x,head.y-tail.y}){
	case point{-2,-2}, point{2,-1},point{1,-2}, point{0,-2}, point{-1,-2}, point{-2,-1}, point{2,-2}:
		t.y--
	}
	switch (point{head.x-tail.x,head.y-tail.y}){
	case point{-2,-2}, point{-1,-2},point{-2,-1}, point{-2,-0}, point{-2,1}, point{-1,2}, point{-2,2}:
		t.x--
	}
	return
}

func main() {
	file, _ := os.Open("file.txt")

	defer file.Close()

	sc := bufio.NewScanner(file)

	visited := make(map[point]bool)
	knots := make([]point, 10)

	visited[knots[9]] = true

	var direction string
	var steps int

	for sc.Scan() {
		fmt.Sscanf(sc.Text(), "%s %d", &direction, &steps)

		for steps > 0 {
			switch direction {
				case "U":
					knots[0].y++
				case "R":
					knots[0].x++
				case "D":
					knots[0].y--
				case "L":
					knots[0].x--
			}
			
			for i := range knots[:len(knots) - 1] {
				knots[i + 1] = adjust(knots[i + 1], knots[i])
			}

			steps--

			visited[knots[9]] = true
		}
	}
	fmt.Println(len(visited))
}