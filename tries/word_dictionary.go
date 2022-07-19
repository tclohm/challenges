package main

import "fmt"

type WordDictionary struct {
	children []*WordDictionary
	isWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		children: make([]*WordDictionary, 26),
		isWord: false,
	}
}

func (this *WordDictionary) AddWord(word string) {
	t := this
	for _, r := range word {
		if t.children[r - 'a'] == nil {
			new := Constructor()
			t.children[r - 'a'] = &new
		}
		t = t.children[r - 'a']
	}
	t.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	t := this
	for _, r := range word {
		if r == '.' {
			for alpha := 0 ; alpha < 26 ; alpha++ {
				if t.children[alpha] != nil && t.children[alpha].Search(word[alpha:]) {
					return true
				}
			}

			return false
		}

		if t.children[r - 'a'] == nil {
			return false
		}

		t = t.children[r - 'a']
	}
	return t.isWord
}

func main() {
	wd := Constructor()

	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	fmt.Println(wd.Search("pad")) // return False
	fmt.Println(wd.Search("bad")) // return True
	fmt.Println(wd.Search(".ad")) // return True
	fmt.Println(wd.Search("b.."))
}