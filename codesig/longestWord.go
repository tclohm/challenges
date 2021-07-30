import "regexp" 

func longestWord(text string) string {
    longest := ""
    re := regexp.MustCompile("[A-Za-z]+")
    for _, elem := range re.FindAllString(text, -1) {
        if len(elem) > len(longest) {
            longest = elem
        }
    }
    return longest
}

import "unicode"

func cleanup(text string) []string {
    result := []string{}
    tmp := ""
    for _, r := range text {

        if unicode.IsLetter(r) {
            tmp += string(r)
        }
        if unicode.IsSpace(r) {
            result = append(result, tmp)
            tmp = ""
        }
    }
    result = append(result, tmp)
    return result
}

func longestWord(text string) string {
    array := cleanup(text)
    big := array[0]
    
    for _, r := range array {
        if len(r) > len(big) {
            big = string(r)
        }
    }
    
    return big
}
