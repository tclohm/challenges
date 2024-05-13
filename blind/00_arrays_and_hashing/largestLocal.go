package main

import (
  "fmt"
)

func largestLocal(grid [][]int) [][]int {
  length := len(grid)
  result := make([][]int, length - 2)

  for i := 0 ; i < len(result) ; i++ {
    result[i] = make([]int, length - 2)
  }

  for i := 1 ; i < length - 1 ; i++ {
    for j := 1 ; j < length - 1 ; j++ {
      tmp := 0

      for k := i - 1 ; k <= i + 1 ; k++ {
        for l := j - 1 ; l <= j + 1 ; l++ {
          tmp = max(tmp, grid[k][l])
        }
      }

      result[i - 1][j - 1] = tmp
    }
  }
  return result
}


func main() {
  fmt.Println(largestLocal([][]int{{9,9,8,1},{5,6,2,6},{8,2,6,4},{6,2,2,2}}))
  fmt.Println(largestLocal([][]int{{1,1,1,1,1},{1,1,1,1,1},{1,1,1,1,1},{1,1,1,1,1},{1,1,1,1,1}}))
}
