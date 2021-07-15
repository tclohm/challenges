import . "strings"
func stringsConstruction(a string, b string) int {
    countMap := make(map[rune]int)
    for _, rn := range a {
        countMap[rn]++
    }
    count := int(^uint(0) >> 1) // max
    for key, value := range countMap {
        min := Count(b, string(key)) / value
        if min < count {
           count = min 
        }
    }
    return count
}
