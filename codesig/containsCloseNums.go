func containsCloseNums(nums []int, k int) bool {
    if k <= 0 {
        return false
    }
    
    ht := map[int]int{}
    for index, num := range nums {
        if idx, exist := ht[num]; exist {
            if index - idx <= k {
                return true
            }
        }
        ht[num] = index
    }
    return false;
}