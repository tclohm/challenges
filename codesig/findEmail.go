import . "unicode"
func findEmailDomain(address string) string {
    for i, r := range address {
        if string(r) == "@" && IsLetter(rune(address[i+1])) {
            return address[i+1:]
        }
    }
    return "null"
}
