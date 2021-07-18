func roadsBuilding(cities int, roads [][]int) [][]int {
    // create array with single city, size of city
    connected := make([][]bool, cities)
    for i := 0; i < cities; i++ {
        connected[i] = make([]bool, cities)
    }
    // look in roads and build connected roads
    for _, v := range roads {
        cityA, cityB := v[0], v[1]
        connected[cityA][cityB] = true
        connected[cityB][cityA] = true
    }
    
    // move through, and check if connected[index][jndex] != true
    // if not append it to the array of roads to be built
    toBeBuilt := [][]int{}
    for i := 0; i < cities - 1; i++ {
        for j := i+1; j < cities; j++ {
            if connected[i][j] != true {
                toBeBuilt = append(toBeBuilt, []int{i, j})
            }
        }
    }
    return toBeBuilt
