func circleOfNumbers(n int, firstNumber int) int {
    halfwaypoint := n / 2
    opposite := n-halfwaypoint+firstNumber
    if opposite >= n {
        opposite = opposite - n
    }
    
    return opposite
}