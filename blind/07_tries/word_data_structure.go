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
    	bite := word[i]
    	if node.children[bite] == nil {
    		node.children[bite] = &TrieNode{children: make(map[byte]*TrieNode)}
    	}
    	node = node.children[bite]
    }
    node.word = true
}


func (this *WordDictionary) Search(word string) bool {
	var dfs func(int, *TrieNode) bool
	dfs = func(i int, root *TrieNode) bool {
		node := root

		for index := i ; index < len(word) ; index++ {
			bite := word[index]

			if bite == '.' {
				for _, child := range node.children {
					// backtracking or recursion
					if dfs(index + 1, child) {
						return true
					}
				}
				return false
			} else {
				if _, ok := node.children[bite]; !ok {
					return false
				}
				node = node.children[bite]
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