import "math"
func isPower(n int) bool {
    if n == 1 {
        return true
    }
    
    for base := 2 ; int(math.Pow(float64(base), float64(2))) <= n ; base++ {
        for exponent := 2 ; int(math.Pow(float64(base), float64(exponent))) <= n ; exponent++ {
            if int(math.Pow(float64(base), float64(exponent))) == n {
                return true
            }
        }
    }
    return false
    
}
