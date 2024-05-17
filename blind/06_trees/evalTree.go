package main

import (
  tree "trees/common"
  "fmt"
)

func evaluate (root tree.Node) bool {
  if nil == root.Left { return root.Value == 1 }

  if root.Value == 2 {
    return evaluate(root.Left) && evaluate(root.Right)
  }

  if root.Value == 3 {
    return evaluate(root.Left) && evaluate(root.Right)
  }
}

func main() {
  fmt.Println(evaluate(root))
}
