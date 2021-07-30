import (
	"math"
)
func deleteDigit(n int) int {
    m := 0
	for i:= 1;i <= n; i*=10 {
		m = int(math.Max(float64(m), float64(((n/i)/10)*i + n%i)))
	}
    return m 
}


func deleteDigit(n int) int {
    max := 0
    s := strconv.Itoa(n)
    for i := 0; i < len(s); i++ {
        tmp := s[:i] + s[i+1:]
        x, _ := strconv.Atoi(string(tmp))
        if x > max {
            max = x
        }
    }
    return max
}