func possibleSums(coins []int, quantity []int) int {
    ht := map[int]bool{}
    ht[0] = true
    for index, coin := range coins {
        tmp := make(map[int]bool)
        for jndex := 1; jndex <= quantity[index]; jndex++ {
            for ps := range ht {
                tmp[ps + coin*jndex] = true
            }
        }
        for t := range tmp {
            ht[t] = true
        }
    }
    return len(ht) - 1
}