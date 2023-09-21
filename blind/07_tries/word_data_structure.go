package main

import (
	"fmt"
)

type TrieNode struct {
	children map[rune]*TrieNode
	word bool
}

type WordDictionary struct {
    root *TrieNode
}


func Constructor() WordDictionary {
    return WordDictionary{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}


func (this *WordDictionary) AddWord(word string)  {
	node := this.root
    for _, r := range word {
    	if node.children[r] == nil {
    		node.children[r] = &TrieNode{children: make(map[rune]*TrieNode)}
    	}
    	node = node.children[r]
    }
    node.word = true
}


func (this *WordDictionary) Search(word string) bool {
	var dfs func(i int, root *TrieNode) bool
	dfs = func(i int, root *TrieNode) bool {
		node := root

		for index := i ; index < len(word) ; index++ {
			r := rune(word[index])

			if r == '.' {
				for _, child := range node.children {
					if dfs(i + 1, child) {
						return true
					}
				}
				return false
			} else {
				if _, ok := node.children[r]; !ok {
					return false
				}
				node = node.children[r]
			}
		}
		return node.word
	}

	return dfs(0, this.root)
}

func main() {
	word_dict := Constructor()
	word_dict.AddWord("apple")
	word_dict.AddWord("banana")
	fmt.Println(word_dict.Search("apple"))
	fmt.Println(word_dict.Search("app.e"))
	fmt.Println(word_dict.Search("b.nana"))
	fmt.Println(word_dict.Search("ba.a.a"))
}