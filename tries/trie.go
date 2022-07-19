package main

import "fmt"

type Trie struct {
	children [26]*Trie
	Word bool
}

func Constructor() Trie {
	return Trie{[26]*Trie{}, false}
}

func (this *Trie) Insert(word string) {
	t := this

	for _, r := range word {
		if t.children[r - 'a'] == nil {
			new := Constructor()
			t.children[r - 'a'] = &new
		}
		t = t.children[r - 'a']
	}
	t.Word = true
}

func (this *Trie) Search(word string) bool {
	t := this
	for _, r := range word {
		if t.children[r - 'a'] == nil {
			return false
		}
		t = t.children[r - 'a']
	}
	return t.Word == true
}

func (this *Trie) StartsWith(prefix string) bool {
	t := this
	for _, r := range prefix {
		if t.children[r - 'a'] == nil {
			return false
		}
		t = t.children[r - 'a']
	}
	return true
}

func main() {
	// use lowercase
	t := Constructor()
	t.Insert("hello")
	isFound := t.Search("hello")
	has := t.StartsWith("he")

	fmt.Println(isFound, has)
}