import . "strings"

func exist(l string, in []string) bool {
    for _, r := range in {
        if l == string(r) {
            return true
        }
    }
    return false
}

func isCryptSolution(crypt []string, solution [][]string) bool {
    first := Split(crypt[0], "")
    second := Split(crypt[1], "")
    answer := Split(crypt[2], "")
    ht := map[string]string{}
    
    for _, codex := range solution {
        letter, number := codex[0], codex[1]
        ht[letter] = number
    }
    
    u1 := ""
    u2 := ""
    uA := ""
    
    if (ht[string(first[0])] == "0" || ht[string(second[0])] == "0" || ht[string(answer[0])] == "0") && (len(first) > 1 && len(second) > 1 && len(answer) > 1) {
        return false
    }
    
    for _, l := range first {
        u1 += ht[l]
    }
    
    for _, l := range second {
        u2 += ht[l]
    }
    
    for _, l := range answer {
        uA += ht[l]
    }
    
    n1, _ := strconv.Atoi(u1)
    n2, _ := strconv.Atoi(u2)
    nA, _ := strconv.Atoi(uA)
    
    return n1 + n2 == nA
}
