package main

import "fmt"
// Time : 2^n, Space : n
func wordBreak(s string, dict []string) []string {
  var backtracking func (index int)
  backtracking = func (index int) {}
  return []string{}
}

func main() {
  fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
  fmt.Println(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
  fmt.Println(wordBreak("catsanddog", []string{"cat", "dog", "sand", "and", "cat"}))
}
