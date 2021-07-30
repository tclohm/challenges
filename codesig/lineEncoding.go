import . "strconv"
func lineEncoding(s string) string {
    build := ""
    head := string(s[0])
    s = s[1:]
    num := 1
    
    for len(s) > 0 {
        if head == string(s[0]) {
            num++
            s = string(s[1:])
        } else if head != string(s[0]) && num > 1 {
            build += Itoa(num) + head
            num = 1
            head = string(s[0])
            s = string(s[1:])
        } else {
            build += head
            head = string(s[0])
            s = string(s[1:])
        }
    }
    if num > 1 {
        return build + Itoa(num) + head
    }
    return build + head
}