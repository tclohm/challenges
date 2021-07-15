func magicalWell(a int, b int, n int) int {
    money := 0
    for n > 0 {
        money += (a * b)
        a++
        b++
        n--
    }
    return money
}
