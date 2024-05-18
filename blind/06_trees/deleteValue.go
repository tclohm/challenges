package main

import (
  "fmt"
  "tree" common/tree
)

func deleteValue(root *tree.Node) *tree.Node {
  if nil == root { return nil }
  root.Left = deleteValue(root.Left, target)
  root.Right = deleteValue(root.Right, target)

  if nil == root.Left && nil == root.Right && root.Value == target {
    return nil
  }
  return root
}

func main() {}
