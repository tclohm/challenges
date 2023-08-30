package main

import "fmt"

func isValidSudoku(board [][]string) bool {
	hm := make(map[string]bool)

	for row := 0 ; row < 9 ; row++ {
		for col := 0 ; col < 9 ; col++ {
			value := string(board[row][col])

			if value == "." { continue }

			_, exist_row := hm[value + "in row" + string(row)]
			_, exist_col := hm[value + "in col" + string(col)]
			_, exist_gri := hm[value + "in grid" + string(row/3) + "-" + string(col/3)]

			if exist_row || exist_col || exist_gri { 
				return false 
			} else {
				hm[value + "in row" + string(row)] = true
				hm[value + "in col" + string(col)] = true
				hm[value + "in grid" + string(row/3) + "-" + string(col/3)] = true
			}
		}
	}
    return true
}

func main() {
	fmt.Println(isValidSudoku([][]string{
		 {"5","3",".",".","7",".",".",".","."},
		{"6",".",".","1","9","5",".",".","."},
		{".","9","8",".",".",".",".","6","."},
		{"8",".",".",".","6",".",".",".","3"},
		{"4",".",".","8",".","3",".",".","1"},
		{"7",".",".",".","2",".",".",".","6"},
		{".","6",".",".",".",".","2","8","."},
		{".",".",".","4","1","9",".",".","5"},
		{".",".",".",".","8",".",".","7","9"},},
	))

	fmt.Println(isValidSudoku([][]string{
		{"8","3",".",".","7",".",".",".","."},
		{"6",".",".","1","9","5",".",".","."},
		{".","9","8",".",".",".",".","6","."},
		{"8",".",".",".","6",".",".",".","3"},
		{"4",".",".","8",".","3",".",".","1"},
		{"7",".",".",".","2",".",".",".","6"},
		{".","6",".",".",".",".","2","8","."},
		{".",".",".","4","1","9",".",".","5"},
		{".",".",".",".","8",".",".","7","9"},},
	))

}