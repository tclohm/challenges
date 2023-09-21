package main

import (
	"fmt"
)

type TrieNode struct {
	children map[byte]*TrieNode
	word bool
}

type WordDictionary struct {
    root *TrieNode
}


func Constructor() WordDictionary {
    return WordDictionary{root: &TrieNode{children: make(map[byte]*TrieNode)}}
}


func (this *WordDictionary) AddWord(word string)  {
	node := this.root
    for i, _ := range word {
    	if node.children[word[i]] == nil {
    		node.children[word[i]] = &TrieNode{children: make(map[byte]*TrieNode)}
    	}
    	node = node.children[word[i]]
    }
    node.word = true
}


func (this *WordDictionary) Search(word string) bool {
	var dfs func(i int, root *TrieNode) bool
	dfs = func(i int, root *TrieNode) bool {
		node := root

		for index := i ; index < len(word) ; index++ {
			b := word[index]

			if b == '.' {
				for _, child := range node.children {
					if dfs(index + 1, child) {
						return true
					}
				}
				return false
			} else {
				if _, ok := node.children[b]; !ok {
					return false
				}
				node = node.children[b]
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