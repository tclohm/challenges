import "sort"

func electionsWinners(votes []int, k int) int {

    count := 0
    sort.Ints(votes)
    max := votes[len(votes)-1]
    if k == 0 {
        z := 0
        for _, v := range votes {
            if v == max {
                z++
            }
        }
        if z > 1 {
            return 0
        }
        return 1
    }
    for _, v := range votes {
        if v + k > max {
            count++
        }
    }
    
    return count
}
