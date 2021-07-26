func digitDegree(n int) int {
    degree := 0
    tmp := 0
    
    if n < 10 {
        return degree
    }
    
    for n != 0 {
        remainder := n % 10
        n /= 10
        tmp += remainder
        
        if n == 0 && tmp > 9 {
            n = tmp
            tmp = 0
            degree++
        }
    }
    
    if tmp > 0 {
        return degree + 1
    }
    return degree
}