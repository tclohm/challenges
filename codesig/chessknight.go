import . "strconv"
func buildboard() [][]string {
    board := [][]string{}
    for i := 0 ; i < 8 ; i++ {
        row := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
        for j := 0 ; j < len(row) ; j++ {
            row[j] = row[j] + Itoa(i + 1)
        }
        board = append(board, row)
    }
    return board
}

func chessKnight(cell string) int {

    board := buildboard()
    
    valid := 0
    
    possibleMoves := [][]int{
        {-2, -1}, 
        {-1, -2}, 
        {2, -1}, 
        {1, -2}, 
        {2, 1}, 
        {1, 2},
        {-1, 2},
        {-2, 1},
    }

    for row := 0 ; row < len(board) ; row++ {
        for col := 0 ; col < len(board[row]) ; col++ {
            if board[row][col] == cell {
                for i := 0 ; i < len(possibleMoves) ; i++ {
                    if row - possibleMoves[i][0] < 8 && row - possibleMoves[i][0] >= 0 && col - possibleMoves[i][1] < 8 && col - possibleMoves[i][1] >= 0 {
                        valid++
                    }
                }
            }
        }
    }
    
    return valid
}
