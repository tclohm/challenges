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

func rotateImage(a[][] int) [][]int {
    n := len(a)
    for i := 0; i < n / 2 ; i++ {
        for j := i; j < n-i-1; j++){
            temp := a[i][j]

            a[i][j] = a[n-j-1][i]

            a[n-j-1][i] = a[n-1-i][n-1-j]

            a[n-1-i][n-1-j] = a[j][n-1-i]

            a[j][n-1-i] = temp

        }
    }
    return a
}