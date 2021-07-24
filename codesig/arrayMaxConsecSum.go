func arrayMaxConsecutiveSum(inputArray []int, k int) int {
    sum := 0
    
    for i := 0; i < k; i++ {
        sum += inputArray[i]
    }
    
    max := sum
    
    for i := 0; i + k < len(inputArray); i++ {
        sum -= inputArray[i]
        sum += inputArray[i+k]
        
        if sum > max {
            max = sum
        }
    }
    
    return max
}