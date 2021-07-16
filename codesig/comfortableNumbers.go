import "math"
func getSum(n int) int {
    sum := 0
    for n != 0 {
        sum += n % 10
        n = int(math.Floor(float64(n/10)))
    }
    return sum
}

func comfortable(i int, j int) bool {
    digits := getSum(i)
    return i != j && (i - digits <= j) && (j <= i + digits)
}

func comfortableNumbers(l int, r int) int {
    count := 0
    for i := l ; i < r ; i++ {
        for j := i+1 ; j <= r ; j++ {
            if comfortable(i,j) && comfortable(j,i) {
                count++
            } else {
                break
            }
        }
    }
    return count
