unc spiralNumbers(n int) [][]int {
    spiral := [][]int{}
    
    for i := 0 ; i < n ; i++ {
        row := make([]int, n, n)
        spiral = append(spiral, row)
    }
    
    top := 0
    bottom := len(spiral) - 1
    left := 0
    right := len(spiral[0]) - 1
    
    num := 1
    
    for left <= right && top <= bottom {
        for i := left ; i <= right ; i++ {
            spiral[top][i] = num
            num++
        }
        top++
        for i := top ; i <= bottom ; i++ {
            spiral[i][right] = num
            num++
        }
        right--
        for i := right ; i >= left ; i-- {
            spiral[bottom][i] = num
            num++
        }
        bottom--
        for i := bottom ; i >= top ; i-- {
            spiral[i][left] = num
            num++
        }
        left++
    }
    
    return spiral
}
