package main

import (
	t "tries/common"
	"fmt"
)

func main() {
	trie := t.Constructor()
	trie.Insert("apple")
	trie.Insert("app")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.Search("ap"))
	fmt.Println(trie.StartsWith("app"))
}