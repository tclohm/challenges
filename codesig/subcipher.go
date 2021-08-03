func exists(letter rune, in map[rune]rune) bool {
    for k, _ := range in {
        if letter == k {
            return true
        }
    }
    return false
}
func isSubstitutionCipher(string1 string, string2 string) bool {
    c1:= map[rune]rune{}
    c2:= map[rune]rune{}
    for i, alpha := range string1 {
        al := rune(string2[i])
        
        // does alpha already exist and does it have a different value
        if exists(alpha, c1) && al != c1[alpha] {
            return false
        }
        
        c1[alpha] = al
        c2[al] = alpha
        
    }
    
    return len(c1) == len(c2)
}
