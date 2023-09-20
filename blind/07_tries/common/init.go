package common

type Trie struct {
    Value *Trie
    Children [26]*Trie
    IsEndOfWord bool
}


func Constructor() Trie {
    return Trie{
    	Value: &Trie{},
    	Children: [26]*Trie{},
    	IsEndOfWord: false,
    }
}


func (this *Trie) Insert(word string)  {
    
    for i, _ := range word {
    	index := word[i] - 'a'
    	if this.Children[index] == nil {
    		this.Children[index] = &Trie{}
    	}
    	this = this.Children[index]
    }

    this.IsEndOfWord = true
}


func (this *Trie) Search(word string) bool {

    for i, _ := range word {
    	index := word[i] - 'a'
    	if this.Children[index] == nil {
    		return false
    	}
    	this = this.Children[index]
    }

    return this.IsEndOfWord
}


func (this *Trie) StartsWith(prefix string) bool {

    for i, _ := range prefix {
    	index := prefix[i] - 'a'
    	if this.Children[index] == nil {
    		return false
    	}
    	this = this.Children[index]
    }
    return true
}