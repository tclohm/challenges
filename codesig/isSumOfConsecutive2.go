func isSumOfConsecutive2(n int) int {
    count := 0
    for start := 1 ; start < n - 1 ; start++ {
        next := start + 1
        sum := start + next
        
        for sum < n {
            next += 1
            sum = sum + next
        }
        
        if sum == n {
            count++
        }
    }
    return count
