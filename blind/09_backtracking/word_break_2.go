package main

import "fmt"

func wordBreak(s string, dict []string) []string {
  return []string{}
}

func main() {
  fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
  fmt.Println(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
  fmt.Println(wordBreak("catsanddog", []string{"cat", "dog", "sand", "and", "cat"}))
}
