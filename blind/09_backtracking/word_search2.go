package main

import "fmt"

type TrieNode struct {
	children map[string]*TrieNode
	isWord bool
}

func Constructor() *TrieNode {
	return &TrieNode{
		children: make(map[string]*TrieNode, 26),
		isWord: false,
	}
}

func (self *TrieNode) addWord(word string) {
	for _, r := range word {
		_, exist := self.children[string(r)]; if !exist {
			self.children[string(r)] = Constructor()
		}
		self = self.children[string(r)]
	}
	self.isWord = true
}

func findWords(board [][]string, words []string) []string {
	root := Constructor()
	for _, word := range words {
		root.addWord(word)
	}

	ROWS, COLS := len(board), len(board[0])

	result, visited := []string{}, make(map[int]bool)

	var backtrack func (row, col int, node *TrieNode, current_word string)
	backtrack = func (row, col int, node *TrieNode, current_word string) {
		if nil == node { return }
		fmt.Println("past", node)
		_, exist := node.children[board[row][col]]
		if row < 0 || col < 0 || row == ROWS || col == COLS || visited[row * col] == true || exist {
			return
		}

		visited[row * col] = true
		node = node.children[board[row][col]]
		current_word += board[row][col]
		fmt.Println(current_word)
		if nil != node && node.isWord {
			result = append(result, current_word)
		}

		backtrack(row - 1, col, node, current_word)
		backtrack(row + 1, col, node, current_word)
		backtrack(row, col - 1, node, current_word)
		backtrack(row, col + 1, node, current_word)

		visited[row * col] = false
	}

	for row := 0 ; row < ROWS ; row++ {
		for col := 0 ; col < COLS ; col++ {
			backtrack(row, col, root, "")
		}
	}

	return result
}

func main() {
	fmt.Println(findWords([][]string{
		{"o","a","a","n"}, 
		{"e","t","a","e"}, 
		{"i","h","k","r"}, 
		{"i","f","l","v"},
	}, []string{"oath","pea","eat","rain"}))

	fmt.Println(findWords([][]string{
		{"a","b"},{"c","d"},
	}, []string{"abcb"}))
}