func isEqual(a1 []int, a2 []int) bool {
    if len(a1) != len(a2) {
        return false
    }
    
    for index, char := range a1 {
        if a2[index] != char {
            return false
        }
    }
    return true
}
func newRoadSystem(roadRegister [][]bool) bool {
    // change all the roads to one-directional
    // [[false, true,  false, false],
    // [false, false, true,  false],
    // [true,  false, false, true ],
    // [false, false, true,  false]]
    // { 0: 1, 1: 2, 2: 0, 2: 3, 3: 2 }
    // question is does each city have the same number
    // each city must have the same number of incoming and outgoing roads
    length := len(roadRegister[0])
    outs := make([]int, length, length)
    ins := make([]int, length, length)
    for y, _ := range roadRegister {
        for x, _ := range roadRegister[y] {
            if roadRegister[y][x] == true {
                ins[x] += 1
                outs[y] += 1
            }
        }
    }
    
    return isEqual(ins, outs)
}