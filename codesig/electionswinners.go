func duplicates(node []int, votes []int) bool {
    index, value := node[0], node[1]
    for i, v := range votes {
        if value == v && i != index {
            return true
        }
    }
    return false
}

func electionsWinners(votes []int, k int) int {
    // array of numbers of votes
    // k = people who havent voted
    // # of candidates that can win the election
    // winner == more votes than others
    // there cannot be more than 1 winner
    curr_lead := []int{0,0}
    
    for i, v := range votes {
        highest := curr_lead[1]
        if highest < v {
            curr_lead[0] = i
            curr_lead[1] = v
        }
    }
    
    highest := curr_lead[1]
    //position := curr_lead[0]
    
    total := 0
    
    for _, v := range votes {
        if v + k > highest {
            total++
        }
    }
    
    if k == 0 && duplicates(curr_lead, votes) {
        return 0
    }

    if total == 0 {
        return 1
    }
    return total
}
