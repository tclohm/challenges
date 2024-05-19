package main

import (
  "fmt"
  tree "common/tree"
)

func distribute(root *tree.Node) int {
  res := 0
  var dfs = func (current *tree.Node) []int
  dfs = func(current *tree.Node) []int {
      if nil == current {
        return []int{0,0}
      }

      left_size, left_coins := dfs(current.Left)
      right_size, right_coins := dfs(current.Right)

      size := 1 + left_size + right_size
      coins := current.Value + left_coins + right_coins
      res += size - coins

      return []int{size, coins}
    }

  dfs(root)
  return res
}

func main() {}
