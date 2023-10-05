package main

import "fmt"

type set map[string][]string

func (s set) add(w1, w2 string) {
	s[w1] = append(s[w1], w2)
}

// return a string of the unique letters in the new alien language sorted in lexicographically
// increasing order
// smaller word comes first: ex, ape, app, apple,
func alien(words []string) string {
	adjancency := make(map[string]set)

	for _, word := range words {
		for _, r := range word {
			s := string(r)
			adjancency[s] = make(set)
		}
	}

	for i := 0 ; i < len(words) - 1 ; i++ {
		w1, w2 := words[i], words[i + 1]
		smaller := min(len(w1), len(w2))
		if w1[:smaller] == w2[:smaller] && len(w1) > len(w2) {
			return ""
		}

		for index := 0 ; index < smaller ; index++ {
			if w1[index] != w2[index] {
				adjancency[string(w1[index])].add(string(w1[index]), string(w2[index]))
				break
			}
		}
	}

	visited := make(map[string]bool)
	result := []string{}

	var dfs func (s string) bool
	dfs = func (s string) bool {
		if visited[s] {
			return visited[s]
		}

		visited[s] = true

		for _, neighbors := range adjancency[s] {
			for _, neighbor := range neighbors {
				if dfs(neighbor) {
					return true
				}
			}
		}

		visited[s] = false
		result = append(result, s)
		return false
	}

	for key, _ := range adjancency {

		if dfs(key) {
			return ""
		}
	}

	fmt.Println(result)

	return "ok"
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// sorted order
	words := []string{"wrt", "wrf", "er", "ett", "rftt"}
	alien(words)
}