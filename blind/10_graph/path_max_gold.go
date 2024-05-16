package main

import "fmt"

func maxGold(grid [][]int) int {
  ROWS, COLS := len(grid), len(grid[0])
  visited := make(map[string]bool)
  
  var dfs func(r, c int, visited map[string]bool) int
  dfs = func(r, c int, visited map[string]bool) int {
    strRC := fmt.Sprintf("%d %d", r, c)
    _, exist := visited[strRC]
    if min(r, c) < 0 || r == ROWS || c == COLS || grid[r][c] == 0 || exist {
      return 0
    }

    visited[strRC] = true
    result := grid[r][c]

    for _, neighbor := range [][]int{{r + 1, c}, {r - 1, c}, {r, c + 1}, {r, c - 1}} {
      nR, nC := neighbor[0], neighbor[1]
      result = max(result, grid[r][c] + dfs(nR, nC, visited))
    }
    delete(visited, strRC)
    return result
  }
  result := 0
  for r := range ROWS {
    for c := range COLS {
      result = max(result, dfs(r, c, visited))
    }
  }
  return result
}

func main() {
  fmt.Println(maxGold([][]int{{0,6,0},{5,8,7},{0,9,0}}))
  fmt.Println(maxGold([][]int{{1,0,7},{2,0,6},{3,4,5},{0,3,0},{9,0,20}}))
}
