package main
import (
    "fmt"
)

func isBeautifulString(inputString string) bool {
    alphabet := [26]rune{}
    for _, r := range inputString {
        c := r - 'a'
        alphabet[c]++
    }

    for i := 1; i < len(alphabet); i++ {
        if alphabet[i] > alphabet[i-1] {
            return false
        }
    }
    return true
}


func main() {
    fmt.Println("bbbaacdafe", isBeautifulString("bbbaacdafe"))
    fmt.Println("aabbb", isBeautifulString("aabbb"))
}