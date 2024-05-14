package main

import (
  "fmt"
  "math"
)

// Intuition : task is to maximize the score of a given
//             binay matrix by flipping rows and cols
//
// Approach :
//             1. flip rows : check if the first col is 0, if so flip it 
//             2. flip cols : count the number of zeros, if the # of zeros is greater than half the total rows, flip the cols
//             3. calc score : bitwise left shift to calc 

func (a ) Len() int           { return len(a) }
func (a ) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ) Less(i, j int) bool { return a[i] < a[j] }

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
