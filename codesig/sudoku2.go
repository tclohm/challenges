unc duplicate(arr []string) bool {
    ht := map[string]int{}
    
    for _, s := range arr {
        ht[s]++
        if s != "." && ht[s] > 1 {
            return true
        }
    }
    
    return false
}

func checkX(row int, grid[][]string) bool {
    ht := map[string]int{}
    for _, col := range grid[row] {
        ht[col]++
        if col != "." && ht[col] > 1 {
            return false
        }
    }
    return true
}

func checkY(col int, grid[][]string) bool {
    ht := map[string]int{}
    for i := 0; i < len(grid) ; i++ {
        c := grid[i][col]
        ht[c]++
        if c != "." && ht[c] > 1 {
            return false
        }
    }
    return true
}

func createSquare(top int, left int, right int, bottom int, grid[][]string) []string {
    square := []string{}
    
    for top < bottom {
        part := []string{}
        for i := left ; i < right ; i++ {
            part = append(part, string(grid[top][i]))
        }
        top++
        square = append(square, part...)
    }
    return square
}

func sudoku2(grid [][]string) bool {
    for x, _ := range grid {
        nodup := checkX(x, grid)
        
        if !nodup {
            fmt.Println("fail at x")
            return false
        }
    }
    
    for y, _ := range grid {
        nodup := checkY(y, grid)
        if !nodup {
            fmt.Println("fail at y")
            return false
        }
    }
    
    top := 0
    bottom := 3
    left := 0
    right := 3
    
    for top < 9 && left < 9 {
        square := createSquare(top, left, right, bottom, grid)
    
        if duplicate(square) {
            return false
        }
        
        left += 3
        right += 3
        
        if left > 6 {
            top += 3
            bottom += 3
            left = 0
            right = 3
        }
        
    }
    
    return true
}
