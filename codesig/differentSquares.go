
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
