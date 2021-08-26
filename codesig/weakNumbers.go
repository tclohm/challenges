func weakNumbers(n int) []int {
    divisors := make(map[int]int)
    for i := 1; i <= n; i++ {
        num := findDivisors(i)
        divisors[i] = num
    }
    
    max := 0
    weakness := make(map[int]int)
    for i := 1; i <= n; i++ {
        num := findWeakness(i, divisors)
        weakness[num]++
        if num > max {
            max = num
        }
    }
    
    return []int{max, weakness[max]}
}

func findDivisors(n int) int {
    c := 1
    for i := 1; i < n; i++ {
        if n % i == 0 {
            c++
        }
    }
    
    return c
}

func findWeakness(n int, divisors map[int]int) int {
    count := 0
    for i := 1; i < n; i++ {
        if divisors[i] > divisors[n] {
            count++
        }
    }
    
    return count
}


