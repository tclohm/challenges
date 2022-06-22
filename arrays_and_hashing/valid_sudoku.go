package main

import (
	"fmt"
	//"strconv"
	//"log"
)
type Set map[string]bool

func in(target string, hm Set) bool {
	for k, _ := range hm {
		if target == k {
			return true
		}
	}
	return false
}

func isValidSudoku(board [][]string) bool {
	cols := make(map[int]Set)
	rows := make(map[int]Set)
	// key = (r / 3, c / 3) 	0,0 0,1 0,2
	//					 		1,0 1,1 1,2
	//				     		2,0 2,1 2,2
	var squares map[[2]int]Set = buildSquares()
	
	for r := 0 ; r < 9 ; r++ {
		for c := 0 ; c < 9 ; c++ {
			current_position := board[r][c]
			if current_position == "." {
				continue
			}

			region := [2]int{r / 3, c / 3}

			if in(current_position, rows[r]) || 
			   in(current_position, cols[c]) ||
			   in(current_position, squares[region]) {
			   	return false
			}

			
			
			if _, exist := cols[c]; !exist {
				cols[c] = make(Set)
			}
			cols[c][current_position] = true

			if _, exist := rows[r]; !exist {
				rows[r] = make(Set)
			}
			rows[r][current_position] = true

			if _, exist := squares[region]; !exist {
				squares[region] = make(Set)
			}
			squares[region][current_position] = true

			// fmt.Println("\tROW:", r, "COL:", c, "region:", region, "VALUE:", current_position)
			// fmt.Println("squares:", squares, "\n")
			// fmt.Println("rows:", rows, "\n")
			// fmt.Println("cols:", cols, "\n")

			
		}
	}
	return true
}

func buildSquares() map[[2]int]Set {
	squares := make(map[[2]int]Set)
	for i := 0 ; i < 3 ; i++ {
		for j := 0 ; j < 3 ; j++ {
			pos := [2]int{i, j}
			squares[pos] = make(Set)
		}
	}
	return squares
}

func main() {
	s := [][]string{
		{"5","3",".",".","7",".",".",".","."},
		{"6",".",".","1","9","5",".",".","."},
		{".","9","8",".",".",".",".","6","."},
		{"8",".",".",".","6",".",".",".","3"},
		{"4",".",".","8",".","3",".",".","1"},
		{"7",".",".",".","2",".",".",".","6"},
		{".","6",".",".",".",".","2","8","."},
		{".",".",".","4","1","9",".",".","5"},
		{".",".",".",".","8",".",".","7","9"},
	}

	s2 := [][]string{
		{"8","3",".",".","7",".",".",".","."},
		{"6",".",".","1","9","5",".",".","."},
		{".","9","8",".",".",".",".","6","."},
		{"8",".",".",".","6",".",".",".","3"},
		{"4",".",".","8",".","3",".",".","1"},
		{"7",".",".",".","2",".",".",".","6"},
		{".","6",".",".",".",".","2","8","."},
		{".",".",".","4","1","9",".",".","5"},
		{".",".",".",".","8",".",".","7","9"},
	}

	fmt.Println(isValidSudoku(s))
	fmt.Println(isValidSudoku(s2))
}