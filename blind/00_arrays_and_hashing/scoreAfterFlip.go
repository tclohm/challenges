package main

import (
  "fmt"
  "math"
)

// input : binary matrix 
//  a move consists of toggling (flipping 0s to 1s and 1s to 0s)
//  all the values in a single row or column
//  each row of the matrix is interpreted as a binary number
//  score is the sum of all these binary numbers
//
// goal : find the max possible score by performing any number of moves
//  (including zero moves) on the matrix

// step-by-step approach
// 1. init max_score
// 2. calculate the initial score of the matrix by interpreting each row
//    as a binary number and summing them
// 3. store the initial max_score
// 4. for each row: 
//      4a. toggle the elements in the current row
//      4b. calculate the score of the modified matrix 
//      4c. update max_score if the new score is higher
//      4d. revert the changes made in step 4a (toggle the elements back)
// 5. for each column:
//      5a. toggle the elements in the current column
//      5b. calculate the score of the modified matrix 
//      5c. update max_score if the new score is higher
//      5d. revert the changes made in step 4a (toggle the elements back)
//
// Take away: -- toggling the row or column effectively flips the binary
//               representation of each number in that row/column.
//               Trying all possilbe combinations of row/column toggles
//               We will find the the max possible score
func matrixScore(grid [][]int) int {
  ROWS, COLS := len(grid), len(grid[0])

  for y := 0 ; y < ROWS ; y++ {
    if grid[y][0] == 1 { continue }
    for x := 0 ; x < COLS ; x++ {
      grid[y][x] = toggle(grid[y][x])
    }
  }

  for x := 0 ; x < COLS ; x++ {
    count := 0
    for y := 0 ; y < ROWS ; y++ {
      if grid[y][x] == 1 {
        count += 1
      }
    }
    if float32(count) >= float32(ROWS)/2 {
      continue
    }
    for y := 0 ; y < ROWS ; y++ {
      grid[y][x] = toggle(grid[y][x])
    }
  }

  maximum := 0
  for i := 0 ; i < ROWS ; i++ {
    maximum += binaryTo(grid[i])
  }

  return maximum
}

func binaryTo(arr []int) int {
  N := len(arr)
  var sum float64 = 0 
  for i := 0 ; i < N ; i++ {
    if arr[i] == 1 {
      sum += math.Pow(float64(2), float64(N - i - 1))
    }
  }
  return int(sum)
}

func toggle(num int) int {
  if num == 0 { return 1 }
  return 0 
}

func main() {
  fmt.Println(matrixScore([][]int{{0,0,1,1},{1,0,1,0},{1,1,0,0}}))
  fmt.Println(matrixScore([][]int{{0}}))
}
