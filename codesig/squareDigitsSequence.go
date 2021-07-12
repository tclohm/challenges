func next(n int) int {
    result := 0
    for n > 0 {
        result += (n % 10) * (n % 10) // n^2 or n**2
        n /= 10
    }
    return result
}

func exists(target int, container []int) bool {
    for _, digit := range container {
        if digit == target {
            return true
        }
    }
    return false
}

func squareDigitsSequence(a0 int) int {
    // 0,1,...n | element == sum of squared digits of the prev
    // sequences ends when we discover an element that has already
    // appeared
    set := make([]int, 0, 100)
    set = append(set, a0)
    current := a0
    for true {
        current = next(current)
        if exists(current, set) {
            break
        }
        set = append(set, current)
    }
    return len(set) + 1
}