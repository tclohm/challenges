func transform(code string) rune {
    // 128,64,32,16,8,4,2,1
    r := 0
    pos := 128
    for _, v := range code {
        if string(v) == "1" {
            r += pos
        }
        pos /= 2
    }
    return rune(r)
}
func messageFromBinaryCode(code string) string {
    result := ""
    for i := 0 ; i < len(code) ; i = i + 8 {
       ascii := code[i:i+8]
       
       r := transform(ascii)
       
       result += string(r)
       
    }
    return result
}