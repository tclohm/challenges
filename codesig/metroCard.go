func metroCard(lastNumberOfDays int) []int {
    if lastNumberOfDays == 31 {
        return []int{28, 30, 31}
    }
    return []int{31}
}
