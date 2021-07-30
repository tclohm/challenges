func chessKnight(cell string) int {
    a := []int{1,2,2,1,-1,-2,-2,-1}
    b := []int{-2,-1,1,2,-2,-1,1,2}
    res := 0
    for i:=0; i<8; i++ {
        if cell[0] + byte(b[i]) >= 'a' && cell[0] + byte(b[i]) <= 'h' && cell[1] + byte(a[i]) >= '1' && cell[1] + byte(a[i]) <= '8' {
            res++
        } 
    }
    return res
}
