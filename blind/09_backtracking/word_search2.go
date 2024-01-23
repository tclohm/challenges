package main

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	isWord bool
}

func (self *TrieNode) addWord(word string) {
	for _, r := range word {
		if self.children[r - 'a'] == nil {
			self.children[r - 'a'] = &TrieNode{}
		}
		self = self.children[r - 'a']
	}
	self.isWord = true
}

func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{}
	for _, word := range words {
		root.addWord(word)
	}

	ROWS, COLS := len(board), len(board[0])

	result, visited := []string{}, make([]bool, ROWS*COLS, ROWS*COLS)

	var backtrack func (row, col int, node *TrieNode, current_word string)
	backtrack = func (row, col int, node *TrieNode, current_word string) {
		if nil == node { return }
		if row < 0 || col < 0 || row == ROWS || col == COLS || node.children[board[row][col] - 'a'] == nil {
			return
		}

		visited[row * col] = true
		node = node.children[board[row][col] - 'a']
		current_word += string(board[row][col])
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
	fmt.Println(findWords([][]byte{
		{'o','a','a','n'}, 
		{'e','t','a','e'}, 
		{'i','h','k','r'}, 
		{'i','f','l','v'},
	}, []string{"oath","pea","eat","rain"}))

	fmt.Println(findWords([][]byte{
		{'a','b'},{'c','d'},
	}, []string{"abcb"}))
}