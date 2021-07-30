import (
    "unicode" 
)
    
func sumUpNumbers(inputString string) int {
    total := 0
    tmp := ""
    
    for i, r := range inputString {
        if unicode.IsDigit(r) {
            tmp += string(r)
        } 
        
        if (unicode.IsSpace(r) || unicode.IsLetter(r)) && len(tmp) > 0 {
            in, _ := strconv.Atoi(tmp)
            total += in
            tmp = ""
        }
        
        if string(r) == "+" || string(r) == "," || string(r) == "-" {
            in, _ := strconv.Atoi(tmp)
            total += in
            tmp = ""
        }
        
        if i == len(inputString) - 1 && (len(tmp) > 0 && tmp != "") {
            in, _ := strconv.Atoi(tmp)
            total += in
            tmp = ""
        }
    }
    return total
}

import "regexp"

func sumUpNumbers(inputString string) int {
    total := 0
    re := regexp.MustCompile("[0-9]+")
    for _, elem := range re.FindAllString(inputString, -1) {
        n, _ := strconv.Atoi(elem)
        total = total + n
    }
    return total
}
