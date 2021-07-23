func extractEachKth(inputArray []int, k int) []int {
    result := make([]int, 0, 20)
    for i := 0 ; i < len(inputArray) ; i++ {
        if (i+1) % k != 0 {
            result = append(result, inputArray[i])
        }
    }
    return result
}