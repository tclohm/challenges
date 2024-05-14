package main

import "fmt"

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
  m, n := len(grid), len(grid[0])
  maxScore := 0

  // initial score
  for _, row := range grid {
    score := 0
    for i := 0 ; i < n ; i++ {
      score = (score << 1) + row[i]
    }
    maxScore += score
  }

  // toggle each row
  for i := 0 ; i < m ; i++ {
    // toggle current row
    toggleRow(grid, i)
    // calculate score with the toggled row
    score := 0
    for _, row := range grid {
      rowScore := 0
      for j := 0 ; j < n ; j++ {
        rowScore = (rowScore << 1) + row[j]
      }
      score += rowScore
    }
    if score > maxScore {
      maxScore = score
    }

    // revert the changes to the current row
    toggleRow(grid, i)
  }

  // toggle each column
  for j := 0 ; j < n ; j++ {
    // toggle current column
    toggleColumn(grid, j)
    // calculate score with the toggled column
    score := 0
    for _, row := range grid {
      rowScore := 0
      for k := 0 ; k < n ; k++ {
        rowScore = (rowScore << 1) + row[k]
      }
      score += rowScore
    }
    if score > maxScore {
      maxScore = score
    }

    toggleColumn(grid, j) 
  }

  return maxScore
}

func toggleRow(grid [][]int, row int) {
  for x := range grid[row] {
    grid[row][x] = 1 - grid[row][x]
  }
}

func toggleColumn(grid [][]int, col int) {
  for y := range grid {
    grid[y][col] = 1 - grid[y][col]
  }
}

func main() {
  fmt.Println(matrixScore([][]int{{0,0,1,1},{1,0,1,0},{1,1,0,0}}))
  fmt.Println(matrixScore([][]int{{0}}))
}
