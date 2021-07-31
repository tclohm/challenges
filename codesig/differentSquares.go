
func isDuplicate(square string, in string) bool {
    for index := 0 ; index < len(in) - 1 ; index = index + len(square) {
        partition := in[index:index+len(square)]
        if square == partition {
            return true
        }
    }
    return false
}
func differentSquares(matrix [][]int) int {
    cache := ""
    for y := 0 ; y < len(matrix)-1; y++ {
        for x := 0 ; x < len(matrix[y])-1 ; x++ {
            left        := strconv.Itoa(matrix[y][x])
            right       := strconv.Itoa(matrix[y][x+1])
            bottomleft  := strconv.Itoa(matrix[y+1][x])
            bottomright := strconv.Itoa(matrix[y+1][x+1])
            
            square := left+right+bottomright+bottomleft
            
            // is square in our array
            alreadyIn := isDuplicate(square, cache)
            if !alreadyIn {
                cache += square
            } 
        }
    }
    
    return len(cache) / 4
}


func differentSquares(matrix [][]int) int {
    squares := map[string]int{}
    for r := 0; r < len(matrix)-1; r++ {
        for c := 0; c < len(matrix[0])-1; c++ {
            cell1 := strconv.Itoa(matrix[r][c])
            cell2 := strconv.Itoa(matrix[r+1][c])
            cell3 := strconv.Itoa(matrix[r][c+1])
            cell4 := strconv.Itoa(matrix[r+1][c+1])
            cell := cell1 + cell2 + cell3 + cell4
            squares[cell] = 1
        }
    }
    keys := []string{}
    for k := range squares {
        keys = append(keys, k)
    }
    return len(keys)
}
