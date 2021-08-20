func findProfession(level int, pos int) string {
    /*
    Level 1: E
    Level 2: ED
    Level 3: EDDE
    Level 4: EDDEDEED
    Level 5: EDDEDEEDDEEDEDDE 
    
    Level input is not necessary because first elements are the same
    The result is based on the count of 1's in binary representation of position-1
    If position is even, then Engineer; Else Doctor
    */
    
    bits := IntegerToBinary(pos-1)
    count := 0
    for _, b := range bits {
        if b == '1' {
            count++
        }
    }
    
    if count % 2 == 0 {
        return "Engineer"
    }
    return "Doctor"
}

func IntegerToBinary(n int) string {
   return strconv.FormatInt(int64(n), 2)
}
