import "math"
func knapsackLight(value1 int, weight1 int, value2 int, weight2 int, maxW int) int {
    if maxW >= weight1 + weight2 {
        return value1 + value2
    }
    
    if maxW >= weight1 && maxW >= weight2 {
        return int(math.Max(float64(value1), float64(value2)))
    }
    
    if maxW >= weight1 {
        return value1
    }
    
    if maxW >= weight2 {
        return value2
    }
    
    return 0
    
