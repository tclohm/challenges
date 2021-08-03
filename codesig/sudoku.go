func duplicates(arr []int) bool {
    ht := map[int]int{}
    for _, n := range arr {
        ht[n]++
    }
    
    if len(ht) != 9 {
        return true
    }
    return false
}

func capture(top int, left int, right int, bottom int, grid [][]int) []int {
    main := []int{}
    for top < bottom {
        arr := []int{}
        for i := left ; i < right ; i++ {
            arr = append(arr, grid[top][i])
        }
        top++
        main = append(main, arr...)
    }
    return main
}

func checkX(row int, grid[][]int) bool {
    ht := map[int]int{}
    for _, col := range grid[row] {
        ht[col]++
    }
    if len(ht) != 9 {
        return false
    }
    return true
}

func checkY(col int, grid [][]int) bool {
    ht := map[int]int{}
    for i := 0 ; i < len(grid) ; i++ {
        c := grid[i][col]
        ht[c]++
    }
    if len(ht) != 9 {
        return false
    }
    return true
}

func sudoku(grid [][]int) bool {
    
    // MARK: check the row
    for i, _ := range grid {
        noDuplicates := checkX(i, grid)
        if !noDuplicates {
            return false
        }
    }
    
    // MARK: check the col down
    for i, _ := range grid {
        noDuplicates := checkY(i, grid)
        if !noDuplicates {
            return false
        }
    }
    
    // MARK: check the 3x3 squares
    top := 0
    bottom := 3
    left := 0
    right := 3
    
    for top < 9 && left < 9 {
        
        // capture the array
        arr := capture(top, left, right, bottom, grid)
        
        if duplicates(arr) {
            return false
        }
        
        // left -> right and then when right
         
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