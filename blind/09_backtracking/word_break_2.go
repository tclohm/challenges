package main

import (
  "fmt"
  "strings"
)
// Time : 2^n, Space : n
func wordBreak(s string, dict []string) []string {
  var set = createSet(dict)
  var current = []string{}
  var result = []string{}

  var backtracking func (index int)
  backtracking = func (index int) {
    if index == len(s) {
      result = append(result, strings.Join(current, " "))
      return
    }
    for jndex := index ; jndex < len(s) ; jndex++ {
      word := s[index:jndex+1]
      if _, exist := set[word]; exist {
        current = append(current, word)
        backtracking(jndex + 1)
        // pop
        current = current[:len(current) - 1]
      }
    }
  }

  backtracking(0)
  return result
}

func createSet(words []string) map[string]bool {
  res := map[string]bool{}
  for _, w := range words {
    res[w] = true
  }
  return res
}

func main() {
  fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
  fmt.Println(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
  fmt.Println(wordBreak("catsanddog", []string{"cat", "dog", "sand", "and", "cat"}))
}
