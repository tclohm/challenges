package main

import "fmt"

func solveThat(str string) int {
	var longestScore int = 0
	var currentCount int = 0
	var currentScore int = 0

	for i := 0 ; i < len(str) ; i++ {

		if str[i] == '(' {
			currentCount += 1
		} else {
			currentCount -= 1
		}

		if currentCount < 0 {
			currentCount = 0
			currentScore = 0
			continue
		}

		if str[i] == ')' {
			currentScore += 2
		}

		if currentScore > longestScore {
			longestScore = currentScore
		}

	}

	return longestScore
}

type Queue struct {
	Box [][]int
}

func (q *Queue) Push(item []int) {
	q.Box = append(q.Box, item)
}

func (q *Queue) Shift() (int, int) {
	item := q.Box[0]
	q.Box = q.Box[1:]
	return item[0], item[1]
}

func (q *Queue) Length() int {
	return len(q.Box)
}


type matrix [][]int

func islandCount(matrix matrix) int {
	// left, up, right, down
	directions := [][]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}

	count := 0

	LAND := 1
	WATER := 0

	q := Queue{}

	for r := 0 ; r < len(matrix) ; r++ {
		for c := 0 ; c < len(matrix[0]) ; c++ {
			if matrix[r][c] == WATER { continue }
			if matrix[r][c] == LAND {
				
				point := []int{r,c}
				q.Push(point)
				count++
				matrix[r][c] = WATER

				for q.Length() != 0 {
					row, col := q.Shift()
					for d := 0 ; d < len(directions) ; d++ {
						current := directions[d]
						nextR := current[0] + row
						nextC := current[1] + col

						if nextR < 0 || nextR >= len(matrix) || nextC < 0 || nextC >= len(matrix) { continue }

						if matrix[nextR][nextC] == LAND {
							q.Push([]int{nextR, nextC})
							matrix[nextR][nextC] = WATER
						}
					}
				}
			}
		}
	}
	return count
}

func clearIsland(matrix matrix, r, c int) {
	if r < 0 || r >= len(matrix) ||
	c < 0 || c >= len(matrix[r]) || 
	matrix[r][c] != 1 {
		return
	}

	matrix[r][c] = 0
	clearIsland(matrix, r + 1, c)
	clearIsland(matrix, r - 1, c)
	clearIsland(matrix, r, c + 1)
	clearIsland(matrix, r, c - 1)

}

func numIslands(matrix matrix) int {
	var islandCount = 0
	for r := 0 ; r < len(matrix) ; r++ {
		for c := 0 ; c < len(matrix[0]) ; c++ {
			var item = matrix[r][c]
			if item == 1 {
				clearIsland(matrix, r, c)
				islandCount++
			}
		}
	}
	return islandCount
}

func main() {
	//example := "(()())()"

	//fmt.Println(solveThat(example))

	mp := [][]int{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	// np := [][]int{
	// 	{1, 1, 1, 1, 0},
	// 	{1, 1, 0, 1, 0},
	// 	{1, 1, 0, 0, 0},
	// 	{0, 0, 0, 0, 0},
	// }

	fmt.Println("island count using BFS with queue:", islandCount(mp))

	//fmt.Println("island count using BFS without queue:", numIslands(np))
}