// func isPalindrome(st string) bool {
//     // if string is 0 or 1 return true
//     if len(st) == 0 || len(st) == 1 {
//         return true
//     }
//     // if first item and last item are not the same return false
//     if st[0] != st[len(st)-1] {
//         return false
//     }
//     // move through the palindrom and check
//     return isPalindrome(st[1 : len(st)-1])
// }

func isPalindrome(st string) bool {
    for i, j := 0, len(st) - 1 ; i < j ; i, j = i+1, j-1 {
        if st[i] != st[j] {
            return false
        }
    }
    return true
}

func buildPalindrome(st string) string {
    // build the tai.
    tail := ""
    for _, char := range st {
        // if st + tail == palindrome return
        if isPalindrome(st + tail) {
            return st + tail
        }
        // else add char to the front of tail
        tail =  string(char) + tail
    }
    if isPalindrome(st + tail) {
        return st + tail
    }
    return ""
}
