func rotateImage(a [][]int) [][]int {
    length := len(a[0])
    for i := 0 ; i < length ; i++ {
        for j := i ; j < length ; j++ {
            tmp := a[i][j]
            a[i][j] = a[j][i]
            a[j][i] = tmp
        }
    }
    
    for i := 0 ; i < length ; i++ {
        for j := 0 ; j < (length / 2) ; j++ {
            tmp := a[i][j]
            a[i][j] = a[i][length - 1 - j]
            a[i][length - 1 - j] = tmp
        }
    }
    
    return a
}