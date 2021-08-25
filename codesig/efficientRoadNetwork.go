import "math"

func efficientRoadNetwork(n int, roads [][]int) bool {
    if n <= 1 {
        return true
    }
    
    // Build graph.
    distances := make([][]int, n)
    
    for i := 0; i < n; i++ {
        distances[i] = make([]int, n)
        for j, _ := range distances[i] {
            distances[i][j] = math.MaxInt32
        }
    }
    
    for _, r := range roads {
        distances[r[0]][r[1]] = 1
        distances[r[1]][r[0]] = 1
    }
    
    // Compute all distances. Floyd-Warshall.
    for k := 0; k < n; k++ {
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                d := distances[i][k] + distances[k][j]
                
                if d < distances[i][j] {
                    distances[i][j] = d
                }
            }
        }
    }
    
    // Check if any distances are greater than 2.
    // Disconnected nodes will be math.MaxInt32.
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if distances[i][j] > 2 {
                return false
            }
        }
    }
    
    return true
}
