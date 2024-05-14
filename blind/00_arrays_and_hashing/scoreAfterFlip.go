package main

import "fmt"

func matrixScore(grid [][]int) int {
  rows, cols := len(grid), len(grid[0])

  for r := 0 ; r < rows ; r++ {
    if grid[r][0] == 0 {
      for c := 0 ; c < cols ; c++ {
        if grid[r][c] == 0 {
          grid[r][c] = 1
        } else {
          grid[r][c] = 0
        }
      }
    }
  }

  for c := 0 ; c < cols ; c++ {
    oneCount := 0
    for r := 0 ; r < rows ; r++ {
      oneCount += grid[r][c]
    }

    if oneCount < rows - oneCount {
      for r := 0 ; r < rows ; r++ {
        if grid[r][c] == 1 {
          grid[r][c] = 0
        } else {
          grid[r][c] = 1
        }
      }
    }
  }

  answer := 0
  for r := 0 ; r < rows ; r++ {
    for c := 0 ; c < cols ; c++ {
      answer += grid[r][c] << (cols - c - 1)
    }
  }

  return answer
}

func main() {
  fmt.Println(matrixScore([][]int{{0,0,1,1},{1,0,1,0},{1,1,0,0}}))
  fmt.Println(matrixScore([][]int{{0}}))
}
