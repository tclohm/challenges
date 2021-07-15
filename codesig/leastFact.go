func leastFactorial(n int) int {
    smallest_factorial := 1
    for i := 1 ; smallest_factorial < n ; i++ {
        smallest_factorial *= i
    }
    
    return smallest_factorial
}
