import . "unicode"
func longestDigitsPrefix(inputString string) string {
    max := 0 
    
    for _, r := range inputString {
        if IsDigit(r) {
            max++
        } else {
            break
        }
    }
    
    return inputString[:max]
}