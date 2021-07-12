import "math"
func pagesNumberingWithInk(current int, numberOfDigits int) int {
    digitsToPrint := int(math.Log10(float64(current)))+1
    for numberOfDigits >= digitsToPrint {
        numberOfDigits -= digitsToPrint
        current++
        digitsToPrint = int(math.Log10(float64(current)))+1
    }
    return current - 1
}
