func isTandemRepeat(inputString string) bool {
    length := len(inputString)
    if length % 2 > 0 { return false }
    length = length / 2
    return inputString[:length] == inputString[length:]
}
