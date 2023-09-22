package main

import "fmt"

type TrieNode struct {
	children map[byte]*TrieNode
	isWord bool
	refs int
}

func (this *TrieNode) AddWord(word string) {
	current := this
	current.refs += 1
	for _, r := range word {
		b := byte(r)
		if _, ok := current.children[b]; !ok {
			current.children[b] = &TrieNode{children: make(map[byte]*TrieNode)}
		}
		current = current.children[b]
		current.refs += 1
	}
	current.isWord = true
}

func (this *TrieNode) removeWord(word string) {
	current := this
	current.refs += 1
	for _, r := range word {
		b := byte(r)
		if _, ok := current.children[b]; !ok {
			current = current.children[b]
			current.refs -= 1
		}
	}
}

func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{children: make(map[byte]*TrieNode)}
	for _, word := range words {
		root.AddWord(word)
	}

	ROWS, COLS := len(board), len(board[0])
	result := make(map[string]bool)
	visit := make(map[int]bool)

	var dfs func(int, int, *TrieNode, string)
	dfs = func(r, c int, node *TrieNode, word string) {

		node_child := node.children[board[r][c]]
		boundaries := r < 0 || r >= ROWS || c < 0 || c >= COLS
		
		if boundaries ||
		node_child == nil ||
		node_child.refs < 1 ||
		visit[r * COLS + c] {
			return
		}

		visit[r * COLS + c] = true
		word += string(board[r][c])

		if node_child.isWord {
			node.isWord = false
			result[word] = true
			root.removeWord(word)
		}

		dfs(r + 1, c,     node, word)
		dfs(r - 1, c,     node, word)
		dfs(r,     c + 1, node, word)
		dfs(r,     c - 1, node, word)
		visit[r * COLS + c] = false
	}

	for r := 0 ; r < ROWS ; r++ {
		for c := 0 ; c < COLS ; c++ {
			dfs(r, c, root, "")
		}
	}

	return list(result)
}

func list(mapping map[string]bool) []string {
	result := make([]string, 0, len(mapping))
	for key, _ := range mapping {
		result = append(result, key)
	}
	return result
}

func main() {
	b1 := [][]byte{{'o','a','a','n'},{'e','t','a','e'},{'i','h','k','r'},{'i','f','l','v'}}
	w1 := []string{"oath","pea","eat","rain"}
	fmt.Println(findWords(b1, w1))
}