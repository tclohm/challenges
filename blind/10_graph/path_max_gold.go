package main

import "fmt"

func maxGold(grid [][]int) int {
  ROWS, COLS := len(grid), len(grid[0])
  visited := make(map[string])
  
  var dfs func(r, c int)
  dfs = func(r, c int) {}

  result := 0
  for r in range ROWS {
    for c in range COLS {
      dfs(r, c)
    }
  }
  return result
}

func main() {
  fmt.Println(maxGold([][]int{{0,6,0},{5,8,7},{0,9,0}}))
  fmt.Println(maxGold([][]int{{1,0,7},{2,0,6},{3,4,5},{0,3,0},{9,0,20}}))
}
