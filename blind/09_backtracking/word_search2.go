package main

import "fmt"

func findWords(board [][]string, words []string) []string {

}

func main() {
	fmt.Println(findWords([][]string{
		{"o","a","a","n"}, 
		{"e","t","a","e"}, 
		{"i","h","k","r"}, 
		{"i","f","l","v"}
	}, []string{"oath","pea","eat","rain"}))

	fmt.Println(findWords([][]string{
		{"a","b"},{"c","d"}
	}, []string{"abcb"}))
}