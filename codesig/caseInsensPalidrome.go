import . "strings"
func reverse(s string) string {
    newstring := ""
    for i := len(s)-1 ; i >= 0 ; i-- {
        newstring += string(s[i])
    }
    return newstring
}
func isCaseInsensitivePalindrome(inputString string) bool {
    rev := reverse(inputString)
    return ToLower(inputString) == ToLower(rev)
}
