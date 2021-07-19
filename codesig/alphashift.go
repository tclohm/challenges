func alphabeticShift(inputString string) string {
    //a,z := 97, 122
    newstring := ""
    for _, char := range inputString {
        if char == 122 { char = 96 }
        charshift := char+1
        newstring += string(charshift)
    }
    return newstring
}
